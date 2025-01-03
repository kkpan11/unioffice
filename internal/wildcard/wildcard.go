//
// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// This is a commercial product and requires a license to operate.
// A trial license can be obtained at https://unidoc.io
//
// DO NOT EDIT: generated by unitwist Go source code obfuscator.
//
// Use of this source code is governed by the UniDoc End User License Agreement
// terms that can be accessed at https://unidoc.io/eula/

package wildcard ;func Match (pattern ,name string )(_ed bool ){if pattern ==""{return name ==pattern ;};if pattern =="\u002a"{return true ;};_a :=make ([]rune ,0,len (name ));_fdd :=make ([]rune ,0,len (pattern ));for _ ,_ag :=range name {_a =append (_a ,_ag );
};for _ ,_age :=range pattern {_fdd =append (_fdd ,_age );};_eg :=false ;return _fg (_a ,_fdd ,_eg );};func _gg (_fe ,_gd []rune ,_gf int )int {for len (_gd )> 0{switch _gd [0]{default:if len (_fe )==0{return -1;};if _fe [0]!=_gd [0]{return _gg (_fe [1:],_gd ,_gf +1);
};case '?':if len (_fe )==0{return -1;};case '*':if len (_fe )==0{return -1;};_ddf :=_gg (_fe ,_gd [1:],_gf );if _ddf !=-1{return _gf ;}else {_ddf =_gg (_fe [1:],_gd ,_gf );if _ddf !=-1{return _gf ;}else {return -1;};};};_fe =_fe [1:];_gd =_gd [1:];};return _gf ;
};func Index (pattern ,name string )(_ef int ){if pattern ==""||pattern =="\u002a"{return 0;};_edc :=make ([]rune ,0,len (name ));_cd :=make ([]rune ,0,len (pattern ));for _ ,_abc :=range name {_edc =append (_edc ,_abc );};for _ ,_ca :=range pattern {_cd =append (_cd ,_ca );
};return _gg (_edc ,_cd ,0);};func MatchSimple (pattern ,name string )bool {if pattern ==""{return name ==pattern ;};if pattern =="\u002a"{return true ;};_e :=make ([]rune ,0,len (name ));_fd :=make ([]rune ,0,len (pattern ));for _ ,_g :=range name {_e =append (_e ,_g );
};for _ ,_gc :=range pattern {_fd =append (_fd ,_gc );};_c :=true ;return _fg (_e ,_fd ,_c );};func _fg (_edf ,_fb []rune ,_bb bool )bool {for len (_fb )> 0{switch _fb [0]{default:if len (_edf )==0||_edf [0]!=_fb [0]{return false ;};case '?':if len (_edf )==0&&!_bb {return false ;
};case '*':return _fg (_edf ,_fb [1:],_bb )||(len (_edf )> 0&&_fg (_edf [1:],_fb ,_bb ));};_edf =_edf [1:];_fb =_fb [1:];};return len (_edf )==0&&len (_fb )==0;};