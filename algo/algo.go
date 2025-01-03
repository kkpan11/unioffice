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

package algo ;import _d "strconv";func RepeatString (s string ,cnt int )string {if cnt <=0{return "";};_fb :=make ([]byte ,len (s )*cnt );_gd :=[]byte (s );for _egd :=0;_egd < cnt ;_egd ++{copy (_fb [_egd :],_gd );};return string (_fb );};

// NaturalLess compares two strings in a human manner so rId2 sorts less than rId10
func NaturalLess (lhs ,rhs string )bool {_g ,_dc :=0,0;for _g < len (lhs )&&_dc < len (rhs ){_gc :=lhs [_g ];_a :=rhs [_dc ];_dg :=_f (_gc );_df :=_f (_a );switch {case _dg &&!_df :return true ;case !_dg &&_df :return false ;case !_dg &&!_df :if _gc !=_a {return _gc < _a ;
};_g ++;_dc ++;default:_dfd :=_g +1;_eg :=_dc +1;for _dfd < len (lhs )&&_f (lhs [_dfd ]){_dfd ++;};for _eg < len (rhs )&&_f (rhs [_eg ]){_eg ++;};_da ,_ :=_d .ParseUint (lhs [_g :_dfd ],10,64);_gcb ,_ :=_d .ParseUint (rhs [_g :_eg ],10,64);if _da !=_gcb {return _da < _gcb ;
};_g =_dfd ;_dc =_eg ;};};return len (lhs )< len (rhs );};func _f (_c byte )bool {return _c >='0'&&_c <='9'};