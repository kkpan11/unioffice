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

package utils ;import (_be "bytes";_d "github.com/unidoc/unioffice/document";_e "github.com/unidoc/unioffice/document/convert";_beg "github.com/unidoc/unipdf/v3/model";);

// GetNumPages will try to get actual document page count by converting the document to a PDF first
// and then get the actual page count from the PDF result.
//
// WARNING: This method is currently in experimental state as the PDF result might have incorrect page count.
func GetNumPages (d *_d .Document )(int ,error ){var _ee _be .Buffer ;_a :=_e .ConvertToPdf (d );if _f :=_a .Write (&_ee );_f !=nil {return 0,_f ;};_ec ,_c :=_beg .NewPdfReader (_be .NewReader (_ee .Bytes ()));if _c !=nil {return 0,_c ;};_cb ,_c :=_ec .GetNumPages ();
if _c !=nil {return 0,_c ;};return _cb ,nil ;};