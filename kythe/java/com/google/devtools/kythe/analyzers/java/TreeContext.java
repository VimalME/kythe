/*
 * Copyright 2017 Google Inc. All rights reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package com.google.devtools.kythe.analyzers.java;

import com.google.common.base.Joiner;
import com.google.devtools.kythe.analyzers.java.SourceText.Positions;
import com.google.devtools.kythe.util.Span;
import com.sun.tools.javac.tree.JCTree;
import com.sun.tools.javac.tree.JCTree.JCClassDecl;
import com.sun.tools.javac.tree.JCTree.JCCompilationUnit;
import com.sun.tools.javac.tree.JCTree.JCMethodDecl;
import java.util.LinkedList;
import java.util.List;

/** Kythe context within a {@link JCTree}. */
class TreeContext {
  private final Positions filePositions;
  private final TreeContext up;
  private final JCTree tree;
  private final Span snippet;
  private JavaNode node;

  public TreeContext(Positions filePositions, JCCompilationUnit topLevel) {
    this(filePositions, null, topLevel, null);
  }

  private TreeContext(Positions filePositions, TreeContext up, JCTree tree, Span snippet) {
    this.filePositions = filePositions;
    this.up = up;
    this.tree = tree;
    this.snippet = snippet;
  }

  public String getSourceName() {
    return filePositions.getFilename();
  }

  /** Path relative to source root. */
  public String getSourcePath() {
    // The URI is absolute and hierarchical, so the path will always begin with a slash (see {@link
    // java.net.URI} for details).  We strip that leading slash to return a relative path.
    return filePositions.getSourceFile().toUri().getPath().substring(1);
  }

  public TreeContext down(JCTree tree) {
    return new TreeContext(filePositions, this, tree, snippet);
  }

  public TreeContext downAsSnippet(JCTree tree) {
    return new TreeContext(filePositions, this, tree, filePositions.getSpan(tree));
  }

  public TreeContext up() {
    return up;
  }

  public JCTree getTree() {
    return tree;
  }

  public JavaNode setNode(JavaNode node) {
    if (this.node != null) {
      throw new IllegalStateException("cannot call setNode() twice on same TreeContext");
    }
    this.node = node;
    return node;
  }

  public JavaNode getNode() {
    return node;
  }

  public TreeContext getMethodParent() {
    TreeContext parent = up();
    while (parent != null && !(parent.getTree() instanceof JCMethodDecl)) {
      parent = parent.up();
    }
    return parent;
  }

  public TreeContext getClassOrMethodParent() {
    TreeContext parent = up();
    while (parent != null
        && !(parent.getTree() instanceof JCMethodDecl || parent.getTree() instanceof JCClassDecl)) {
      parent = parent.up();
    }
    return parent;
  }

  public Span getTreeSpan() {
    return filePositions.getSpan(tree);
  }

  public Span getSnippet() {
    return snippet;
  }

  @Override
  public String toString() {
    List<String> parts = new LinkedList<>();
    parts.add("JCTree = " + tree);
    if (node != null) {
      parts.add("JavaNode = " + node);
    }
    parts.add("TreeSpan = " + getTreeSpan());
    if (snippet != null) {
      parts.add("Snippet span = " + snippet);
    }
    List<String> parents = new LinkedList<>();
    TreeContext cur = up;
    while (cur != null) {
      String parent = "" + cur.getTree().getTag();
      String name = NameVisitor.getName(cur.getTree());
      if (name != null) {
        parent += "[" + name + "]";
      }
      parents.add(parent);
      cur = cur.up();
    }
    if (!parents.isEmpty()) {
      parts.add("JCTree parents = " + parents);
    }
    return String.format("TreeContext{\n  %s,\n}", Joiner.on(",\n  ").join(parts));
  }
}
