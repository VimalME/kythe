//- @pkg ref Package
//- Package.node/kind package
package pkg;

// Checks that classes are record/class nodes and enums are sum/enumClass nodes

//- @Classes defines N
//- N.node/kind record
//- N.subkind class
//- N childof Package
public class Classes {

  //- DefaultCtor childof N
  //- DefaultCtor.node/kind function
  //- DefaultCtor typed DefaultCtorType
  //- DefaultCtorType param.0 FnBuiltin
  //- DefaultCtorType param.1 N
  //- !{ DefaultCtorAnchor defines DefaultCtor }

  //- @StaticInner defines SI
  //- SI.node/kind record
  //- SI.subkind class
  //- SI childof N
  private static class StaticInner {
    //- Ctor childof SI
    //- Ctor.node/kind function
    //- Ctor typed CtorType
    //- CtorType param.0 FnBuiltin
    //- CtorType param.1 SI
    //- @StaticInner defines Ctor
    public StaticInner() {}
  }

  //- @Inner defines I
  //- I.node/kind record
  //- I.subkind class
  //- I childof N
  private class Inner {}

  //- @Enum defines E
  //- E.node/kind sum
  //- E.subkind enumClass
  //- E childof N
  private static enum Enum {}
}
