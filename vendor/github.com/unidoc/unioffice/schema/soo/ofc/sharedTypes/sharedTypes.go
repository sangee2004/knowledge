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

package sharedTypes ;import (_e "encoding/xml";_g "fmt";_d "regexp";);var ST_FixedPercentagePatternRe =_d .MustCompile (ST_FixedPercentagePattern );const (ST_TrueFalseUnset ST_TrueFalse =0;ST_TrueFalseT ST_TrueFalse =1;ST_TrueFalseF ST_TrueFalse =2;ST_TrueFalseTrue ST_TrueFalse =3;
ST_TrueFalseFalse ST_TrueFalse =4;);func (_ecgf *ST_ConformanceClass )UnmarshalXML (d *_e .Decoder ,start _e .StartElement )error {_ccga ,_bbd :=d .Token ();if _bbd !=nil {return _bbd ;};if _cbg ,_ad :=_ccga .(_e .EndElement );_ad &&_cbg .Name ==start .Name {*_ecgf =1;
return nil ;};if _bg ,_gcgba :=_ccga .(_e .CharData );!_gcgba {return _g .Errorf ("\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0063\u0068a\u0072\u0020\u0064\u0061\u0074\u0061\u002c\u0020\u0067\u006ft\u0020\u0025\u0054",_ccga );}else {switch string (_bg ){case "":*_ecgf =0;
case "\u0073\u0074\u0072\u0069\u0063\u0074":*_ecgf =1;case "\u0074\u0072\u0061n\u0073\u0069\u0074\u0069\u006f\u006e\u0061\u006c":*_ecgf =2;};};_ccga ,_bbd =d .Token ();if _bbd !=nil {return _bbd ;};if _ceb ,_eda :=_ccga .(_e .EndElement );_eda &&_ceb .Name ==start .Name {return nil ;
};return _g .Errorf ("\u0065\u0078\u0070\u0065c\u0074\u0065\u0064\u0020\u0065\u006e\u0064\u0020\u0065\u006ce\u006de\u006e\u0074\u002c\u0020\u0067\u006f\u0074 \u0025\u0076",_ccga );};func (_bd ST_TrueFalseBlank )Validate ()error {return _bd .ValidateWithPath ("")};
func (_bcd ST_TrueFalse )ValidateWithPath (path string )error {switch _bcd {case 0,1,2,3,4:default:return _g .Errorf ("\u0025s\u003a\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072a\u006eg\u0065 \u0076\u0061\u006c\u0075\u0065\u0020\u0025d",path ,int (_bcd ));
};return nil ;};func (_ca ST_AlgClass )MarshalXMLAttr (name _e .Name )(_e .Attr ,error ){_af :=_e .Attr {};_af .Name =name ;switch _ca {case ST_AlgClassUnset :_af .Value ="";case ST_AlgClassHash :_af .Value ="\u0068\u0061\u0073\u0068";case ST_AlgClassCustom :_af .Value ="\u0063\u0075\u0073\u0074\u006f\u006d";
};return _af ,nil ;};func (_fcd ST_TrueFalse )Validate ()error {return _fcd .ValidateWithPath ("")};func (_efg *ST_VerticalAlignRun )UnmarshalXMLAttr (attr _e .Attr )error {switch attr .Value {case "":*_efg =0;case "\u0062\u0061\u0073\u0065\u006c\u0069\u006e\u0065":*_efg =1;
case "s\u0075\u0070\u0065\u0072\u0073\u0063\u0072\u0069\u0070\u0074":*_efg =2;case "\u0073u\u0062\u0073\u0063\u0072\u0069\u0070t":*_efg =3;};return nil ;};type ST_VerticalAlignRun byte ;func (_dec ST_XAlign )String ()string {switch _dec {case 0:return "";
case 1:return "\u006c\u0065\u0066\u0074";case 2:return "\u0063\u0065\u006e\u0074\u0065\u0072";case 3:return "\u0072\u0069\u0067h\u0074";case 4:return "\u0069\u006e\u0073\u0069\u0064\u0065";case 5:return "\u006fu\u0074\u0073\u0069\u0064\u0065";};return "";
};func (_gbf ST_OnOff1 )MarshalXML (e *_e .Encoder ,start _e .StartElement )error {return e .EncodeElement (_gbf .String (),start );};func (_dad *ST_YAlign )UnmarshalXMLAttr (attr _e .Attr )error {switch attr .Value {case "":*_dad =0;case "\u0069\u006e\u006c\u0069\u006e\u0065":*_dad =1;
case "\u0074\u006f\u0070":*_dad =2;case "\u0063\u0065\u006e\u0074\u0065\u0072":*_dad =3;case "\u0062\u006f\u0074\u0074\u006f\u006d":*_dad =4;case "\u0069\u006e\u0073\u0069\u0064\u0065":*_dad =5;case "\u006fu\u0074\u0073\u0069\u0064\u0065":*_dad =6;};return nil ;
};func (_egc *ST_AlgClass )UnmarshalXML (d *_e .Decoder ,start _e .StartElement )error {_ba ,_ed :=d .Token ();if _ed !=nil {return _ed ;};if _dba ,_bcc :=_ba .(_e .EndElement );_bcc &&_dba .Name ==start .Name {*_egc =1;return nil ;};if _egb ,_ag :=_ba .(_e .CharData );
!_ag {return _g .Errorf ("\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0063\u0068a\u0072\u0020\u0064\u0061\u0074\u0061\u002c\u0020\u0067\u006ft\u0020\u0025\u0054",_ba );}else {switch string (_egb ){case "":*_egc =0;case "\u0068\u0061\u0073\u0068":*_egc =1;
case "\u0063\u0075\u0073\u0074\u006f\u006d":*_egc =2;};};_ba ,_ed =d .Token ();if _ed !=nil {return _ed ;};if _cd ,_fac :=_ba .(_e .EndElement );_fac &&_cd .Name ==start .Name {return nil ;};return _g .Errorf ("\u0065\u0078\u0070\u0065c\u0074\u0065\u0064\u0020\u0065\u006e\u0064\u0020\u0065\u006ce\u006de\u006e\u0074\u002c\u0020\u0067\u006f\u0074 \u0025\u0076",_ba );
};const ST_UniversalMeasurePattern ="\u002d\u003f\u005b\u0030\u002d\u0039\u005d\u002b\u0028\u005c\u002e\u005b\u0030\u002d\u0039\u005d\u002b\u0029\u003f\u0028\u006d\u006d\u007c\u0063m\u007c\u0069\u006e\u007c\u0070t\u007c\u0070c\u007c\u0070\u0069\u0029";
var ST_PositiveUniversalMeasurePatternRe =_d .MustCompile (ST_PositiveUniversalMeasurePattern );const (ST_OnOff1Unset ST_OnOff1 =0;ST_OnOff1On ST_OnOff1 =1;ST_OnOff1Off ST_OnOff1 =2;);func (_bf ST_AlgClass )String ()string {switch _bf {case 0:return "";
case 1:return "\u0068\u0061\u0073\u0068";case 2:return "\u0063\u0075\u0073\u0074\u006f\u006d";};return "";};func (_cc ST_CalendarType )String ()string {switch _cc {case 0:return "";case 1:return "\u0067r\u0065\u0067\u006f\u0072\u0069\u0061n";case 2:return "g\u0072\u0065\u0067\u006f\u0072\u0069\u0061\u006e\u0055\u0073";
case 3:return "\u0067\u0072\u0065\u0067\u006f\u0072\u0069\u0061\u006e\u004d\u0065\u0046r\u0065\u006e\u0063\u0068";case 4:return "\u0067r\u0065g\u006f\u0072\u0069\u0061\u006e\u0041\u0072\u0061\u0062\u0069\u0063";case 5:return "\u0068\u0069\u006ar\u0069";
case 6:return "\u0068\u0065\u0062\u0072\u0065\u0077";case 7:return "\u0074\u0061\u0069\u0077\u0061\u006e";case 8:return "\u006a\u0061\u0070a\u006e";case 9:return "\u0074\u0068\u0061\u0069";case 10:return "\u006b\u006f\u0072e\u0061";case 11:return "\u0073\u0061\u006b\u0061";
case 12:return "g\u0072e\u0067\u006f\u0072\u0069\u0061\u006e\u0058\u006ci\u0074\u0045\u006e\u0067li\u0073\u0068";case 13:return "\u0067\u0072\u0065\u0067or\u0069\u0061\u006e\u0058\u006c\u0069\u0074\u0046\u0072\u0065\u006e\u0063\u0068";case 14:return "\u006e\u006f\u006e\u0065";
};return "";};func (_bfe *ST_TrueFalseBlank )UnmarshalXMLAttr (attr _e .Attr )error {switch attr .Value {case "":*_bfe =0;case "\u0074":*_bfe =1;case "\u0066":*_bfe =2;case "\u0074\u0072\u0075\u0065":*_bfe =3;case "\u0066\u0061\u006cs\u0065":*_bfe =4;case "\u0054\u0072\u0075\u0065":*_bfe =6;
case "\u0046\u0061\u006cs\u0065":*_bfe =7;};return nil ;};func (_eee ST_VerticalAlignRun )MarshalXMLAttr (name _e .Name )(_e .Attr ,error ){_faa :=_e .Attr {};_faa .Name =name ;switch _eee {case ST_VerticalAlignRunUnset :_faa .Value ="";case ST_VerticalAlignRunBaseline :_faa .Value ="\u0062\u0061\u0073\u0065\u006c\u0069\u006e\u0065";
case ST_VerticalAlignRunSuperscript :_faa .Value ="s\u0075\u0070\u0065\u0072\u0073\u0063\u0072\u0069\u0070\u0074";case ST_VerticalAlignRunSubscript :_faa .Value ="\u0073u\u0062\u0073\u0063\u0072\u0069\u0070t";};return _faa ,nil ;};const (ST_XAlignUnset ST_XAlign =0;
ST_XAlignLeft ST_XAlign =1;ST_XAlignCenter ST_XAlign =2;ST_XAlignRight ST_XAlign =3;ST_XAlignInside ST_XAlign =4;ST_XAlignOutside ST_XAlign =5;);const (ST_AlgClassUnset ST_AlgClass =0;ST_AlgClassHash ST_AlgClass =1;ST_AlgClassCustom ST_AlgClass =2;);func (_cdba *ST_AlgType )UnmarshalXMLAttr (attr _e .Attr )error {switch attr .Value {case "":*_cdba =0;
case "\u0074y\u0070\u0065\u0041\u006e\u0079":*_cdba =1;case "\u0063\u0075\u0073\u0074\u006f\u006d":*_cdba =2;};return nil ;};func (_bda ST_TrueFalseBlank )ValidateWithPath (path string )error {switch _bda {case 0,1,2,3,4,6,7:default:return _g .Errorf ("\u0025s\u003a\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072a\u006eg\u0065 \u0076\u0061\u006c\u0075\u0065\u0020\u0025d",path ,int (_bda ));
};return nil ;};func (_gf *ST_TwipsMeasure )Validate ()error {return _gf .ValidateWithPath ("")};func (_fg ST_CalendarType )MarshalXMLAttr (name _e .Name )(_e .Attr ,error ){_fe :=_e .Attr {};_fe .Name =name ;switch _fg {case ST_CalendarTypeUnset :_fe .Value ="";
case ST_CalendarTypeGregorian :_fe .Value ="\u0067r\u0065\u0067\u006f\u0072\u0069\u0061n";case ST_CalendarTypeGregorianUs :_fe .Value ="g\u0072\u0065\u0067\u006f\u0072\u0069\u0061\u006e\u0055\u0073";case ST_CalendarTypeGregorianMeFrench :_fe .Value ="\u0067\u0072\u0065\u0067\u006f\u0072\u0069\u0061\u006e\u004d\u0065\u0046r\u0065\u006e\u0063\u0068";
case ST_CalendarTypeGregorianArabic :_fe .Value ="\u0067r\u0065g\u006f\u0072\u0069\u0061\u006e\u0041\u0072\u0061\u0062\u0069\u0063";case ST_CalendarTypeHijri :_fe .Value ="\u0068\u0069\u006ar\u0069";case ST_CalendarTypeHebrew :_fe .Value ="\u0068\u0065\u0062\u0072\u0065\u0077";
case ST_CalendarTypeTaiwan :_fe .Value ="\u0074\u0061\u0069\u0077\u0061\u006e";case ST_CalendarTypeJapan :_fe .Value ="\u006a\u0061\u0070a\u006e";case ST_CalendarTypeThai :_fe .Value ="\u0074\u0068\u0061\u0069";case ST_CalendarTypeKorea :_fe .Value ="\u006b\u006f\u0072e\u0061";
case ST_CalendarTypeSaka :_fe .Value ="\u0073\u0061\u006b\u0061";case ST_CalendarTypeGregorianXlitEnglish :_fe .Value ="g\u0072e\u0067\u006f\u0072\u0069\u0061\u006e\u0058\u006ci\u0074\u0045\u006e\u0067li\u0073\u0068";case ST_CalendarTypeGregorianXlitFrench :_fe .Value ="\u0067\u0072\u0065\u0067or\u0069\u0061\u006e\u0058\u006c\u0069\u0074\u0046\u0072\u0065\u006e\u0063\u0068";
case ST_CalendarTypeNone :_fe .Value ="\u006e\u006f\u006e\u0065";};return _fe ,nil ;};

// ST_TwipsMeasure is a union type
type ST_TwipsMeasure struct{ST_UnsignedDecimalNumber *uint64 ;ST_PositiveUniversalMeasure *string ;};func (_fb ST_OnOff1 )String ()string {switch _fb {case 0:return "";case 1:return "\u006f\u006e";case 2:return "\u006f\u0066\u0066";};return "";};func (_aag ST_ConformanceClass )ValidateWithPath (path string )error {switch _aag {case 0,1,2:default:return _g .Errorf ("\u0025s\u003a\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072a\u006eg\u0065 \u0076\u0061\u006c\u0075\u0065\u0020\u0025d",path ,int (_aag ));
};return nil ;};type ST_TrueFalse byte ;type ST_AlgType byte ;const (ST_CryptProvUnset ST_CryptProv =0;ST_CryptProvRsaAES ST_CryptProv =1;ST_CryptProvRsaFull ST_CryptProv =2;ST_CryptProvCustom ST_CryptProv =3;);const (ST_YAlignUnset ST_YAlign =0;ST_YAlignInline ST_YAlign =1;
ST_YAlignTop ST_YAlign =2;ST_YAlignCenter ST_YAlign =3;ST_YAlignBottom ST_YAlign =4;ST_YAlignInside ST_YAlign =5;ST_YAlignOutside ST_YAlign =6;);type ST_AlgClass byte ;func (_cfc ST_YAlign )String ()string {switch _cfc {case 0:return "";case 1:return "\u0069\u006e\u006c\u0069\u006e\u0065";
case 2:return "\u0074\u006f\u0070";case 3:return "\u0063\u0065\u006e\u0074\u0065\u0072";case 4:return "\u0062\u006f\u0074\u0074\u006f\u006d";case 5:return "\u0069\u006e\u0073\u0069\u0064\u0065";case 6:return "\u006fu\u0074\u0073\u0069\u0064\u0065";};return "";
};func (_ga ST_AlgClass )MarshalXML (e *_e .Encoder ,start _e .StartElement )error {return e .EncodeElement (_ga .String (),start );};func (_gbd ST_CryptProv )ValidateWithPath (path string )error {switch _gbd {case 0,1,2,3:default:return _g .Errorf ("\u0025s\u003a\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072a\u006eg\u0065 \u0076\u0061\u006c\u0075\u0065\u0020\u0025d",path ,int (_gbd ));
};return nil ;};func (_ceg *ST_TrueFalse )UnmarshalXML (d *_e .Decoder ,start _e .StartElement )error {_fbd ,_fgd :=d .Token ();if _fgd !=nil {return _fgd ;};if _gda ,_cbe :=_fbd .(_e .EndElement );_cbe &&_gda .Name ==start .Name {*_ceg =1;return nil ;
};if _cfgd ,_ec :=_fbd .(_e .CharData );!_ec {return _g .Errorf ("\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0063\u0068a\u0072\u0020\u0064\u0061\u0074\u0061\u002c\u0020\u0067\u006ft\u0020\u0025\u0054",_fbd );}else {switch string (_cfgd ){case "":*_ceg =0;
case "\u0074":*_ceg =1;case "\u0066":*_ceg =2;case "\u0074\u0072\u0075\u0065":*_ceg =3;case "\u0066\u0061\u006cs\u0065":*_ceg =4;};};_fbd ,_fgd =d .Token ();if _fgd !=nil {return _fgd ;};if _ddb ,_bac :=_fbd .(_e .EndElement );_bac &&_ddb .Name ==start .Name {return nil ;
};return _g .Errorf ("\u0065\u0078\u0070\u0065c\u0074\u0065\u0064\u0020\u0065\u006e\u0064\u0020\u0065\u006ce\u006de\u006e\u0074\u002c\u0020\u0067\u006f\u0074 \u0025\u0076",_fbd );};func (_b ST_OnOff )MarshalXML (e *_e .Encoder ,start _e .StartElement )error {e .EncodeToken (start );
if _b .Bool !=nil {e .EncodeToken (_e .CharData (_g .Sprintf ("\u0025\u0064",_gdd (*_b .Bool ))));};if _b .ST_OnOff1 !=ST_OnOff1Unset {e .EncodeToken (_e .CharData (_b .ST_OnOff1 .String ()));};return e .EncodeToken (_e .EndElement {Name :start .Name });
};func (_cegd *ST_VerticalAlignRun )UnmarshalXML (d *_e .Decoder ,start _e .StartElement )error {_fdgc ,_eca :=d .Token ();if _eca !=nil {return _eca ;};if _dbb ,_aba :=_fdgc .(_e .EndElement );_aba &&_dbb .Name ==start .Name {*_cegd =1;return nil ;};if _ffd ,_eec :=_fdgc .(_e .CharData );
!_eec {return _g .Errorf ("\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0063\u0068a\u0072\u0020\u0064\u0061\u0074\u0061\u002c\u0020\u0067\u006ft\u0020\u0025\u0054",_fdgc );}else {switch string (_ffd ){case "":*_cegd =0;case "\u0062\u0061\u0073\u0065\u006c\u0069\u006e\u0065":*_cegd =1;
case "s\u0075\u0070\u0065\u0072\u0073\u0063\u0072\u0069\u0070\u0074":*_cegd =2;case "\u0073u\u0062\u0073\u0063\u0072\u0069\u0070t":*_cegd =3;};};_fdgc ,_eca =d .Token ();if _eca !=nil {return _eca ;};if _agb ,_gec :=_fdgc .(_e .EndElement );_gec &&_agb .Name ==start .Name {return nil ;
};return _g .Errorf ("\u0065\u0078\u0070\u0065c\u0074\u0065\u0064\u0020\u0065\u006e\u0064\u0020\u0065\u006ce\u006de\u006e\u0074\u002c\u0020\u0067\u006f\u0074 \u0025\u0076",_fdgc );};const ST_PercentagePattern ="-\u003f[\u0030\u002d\u0039\u005d\u002b\u0028\u005c\u002e[\u0030\u002d\u0039\u005d+)\u003f\u0025";
func (_bb ST_TwipsMeasure )MarshalXML (e *_e .Encoder ,start _e .StartElement )error {e .EncodeToken (start );if _bb .ST_UnsignedDecimalNumber !=nil {e .EncodeToken (_e .CharData (_g .Sprintf ("\u0025\u0064",*_bb .ST_UnsignedDecimalNumber )));};if _bb .ST_PositiveUniversalMeasure !=nil {e .EncodeToken (_e .CharData (*_bb .ST_PositiveUniversalMeasure ));
};return e .EncodeToken (_e .EndElement {Name :start .Name });};func (_gcgb ST_YAlign )Validate ()error {return _gcgb .ValidateWithPath ("")};func (_fff ST_VerticalAlignRun )String ()string {switch _fff {case 0:return "";case 1:return "\u0062\u0061\u0073\u0065\u006c\u0069\u006e\u0065";
case 2:return "s\u0075\u0070\u0065\u0072\u0073\u0063\u0072\u0069\u0070\u0074";case 3:return "\u0073u\u0062\u0073\u0063\u0072\u0069\u0070t";};return "";};func (_fge ST_AlgClass )Validate ()error {return _fge .ValidateWithPath ("")};func (_c *ST_TwipsMeasure )ValidateWithPath (path string )error {_gc :=[]string {};
if _c .ST_UnsignedDecimalNumber !=nil {_gc =append (_gc ,"\u0053T\u005f\u0055\u006e\u0073\u0069\u0067\u006e\u0065\u0064\u0044\u0065c\u0069\u006d\u0061\u006c\u004e\u0075\u006d\u0062\u0065\u0072");};if _c .ST_PositiveUniversalMeasure !=nil {_gc =append (_gc ,"S\u0054\u005f\u0050\u006f\u0073\u0069t\u0069\u0076\u0065\u0055\u006e\u0069\u0076\u0065\u0072s\u0061\u006c\u004de\u0061s\u0075\u0072\u0065");
};if len (_gc )> 1{return _g .Errorf ("%\u0073\u0020\u0074\u006f\u006f\u0020m\u0061\u006e\u0079\u0020\u006d\u0065\u006d\u0062\u0065r\u0073\u0020\u0073e\u0074:\u0020\u0025\u0076",path ,_gc );};return nil ;};func (_aed ST_XAlign )MarshalXMLAttr (name _e .Name )(_e .Attr ,error ){_eece :=_e .Attr {};
_eece .Name =name ;switch _aed {case ST_XAlignUnset :_eece .Value ="";case ST_XAlignLeft :_eece .Value ="\u006c\u0065\u0066\u0074";case ST_XAlignCenter :_eece .Value ="\u0063\u0065\u006e\u0074\u0065\u0072";case ST_XAlignRight :_eece .Value ="\u0072\u0069\u0067h\u0074";
case ST_XAlignInside :_eece .Value ="\u0069\u006e\u0073\u0069\u0064\u0065";case ST_XAlignOutside :_eece .Value ="\u006fu\u0074\u0073\u0069\u0064\u0065";};return _eece ,nil ;};type ST_YAlign byte ;func (_ccc ST_TrueFalseBlank )String ()string {switch _ccc {case 0:return "";
case 1:return "\u0074";case 2:return "\u0066";case 3:return "\u0074\u0072\u0075\u0065";case 4:return "\u0066\u0061\u006cs\u0065";case 6:return "\u0054\u0072\u0075\u0065";case 7:return "\u0046\u0061\u006cs\u0065";};return "";};type ST_XAlign byte ;func (_cae ST_XAlign )Validate ()error {return _cae .ValidateWithPath ("")};
var ST_PositivePercentagePatternRe =_d .MustCompile (ST_PositivePercentagePattern );func (_aea ST_XAlign )ValidateWithPath (path string )error {switch _aea {case 0,1,2,3,4,5:default:return _g .Errorf ("\u0025s\u003a\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072a\u006eg\u0065 \u0076\u0061\u006c\u0075\u0065\u0020\u0025d",path ,int (_aea ));
};return nil ;};func (_a *ST_OnOff )Validate ()error {return _a .ValidateWithPath ("")};func (_dc ST_OnOff1 )ValidateWithPath (path string )error {switch _dc {case 0,1,2:default:return _g .Errorf ("\u0025s\u003a\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072a\u006eg\u0065 \u0076\u0061\u006c\u0075\u0065\u0020\u0025d",path ,int (_dc ));
};return nil ;};func (_fad ST_AlgType )ValidateWithPath (path string )error {switch _fad {case 0,1,2:default:return _g .Errorf ("\u0025s\u003a\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072a\u006eg\u0065 \u0076\u0061\u006c\u0075\u0065\u0020\u0025d",path ,int (_fad ));
};return nil ;};const ST_PositivePercentagePattern ="\u005b0\u002d9\u005d\u002b\u0028\u005c\u002e[\u0030\u002d9\u005d\u002b\u0029\u003f\u0025";const (ST_VerticalAlignRunUnset ST_VerticalAlignRun =0;ST_VerticalAlignRunBaseline ST_VerticalAlignRun =1;ST_VerticalAlignRunSuperscript ST_VerticalAlignRun =2;
ST_VerticalAlignRunSubscript ST_VerticalAlignRun =3;);func ParseUnionST_OnOff (s string )(ST_OnOff ,error ){_gca :=ST_OnOff {};switch s {case "\u0074\u0072\u0075\u0065","\u0031","\u006f\u006e":_dd :=true ;_gca .Bool =&_dd ;default:_df :=false ;_gca .Bool =&_df ;
};return _gca ,nil ;};

// ST_OnOff is a union type
type ST_OnOff struct{Bool *bool ;ST_OnOff1 ST_OnOff1 ;};func (_abc ST_AlgType )MarshalXML (e *_e .Encoder ,start _e .StartElement )error {return e .EncodeElement (_abc .String (),start );};func (_aca *ST_ConformanceClass )UnmarshalXMLAttr (attr _e .Attr )error {switch attr .Value {case "":*_aca =0;
case "\u0073\u0074\u0072\u0069\u0063\u0074":*_aca =1;case "\u0074\u0072\u0061n\u0073\u0069\u0074\u0069\u006f\u006e\u0061\u006c":*_aca =2;};return nil ;};func (_ac ST_AlgType )MarshalXMLAttr (name _e .Name )(_e .Attr ,error ){_fee :=_e .Attr {};_fee .Name =name ;
switch _ac {case ST_AlgTypeUnset :_fee .Value ="";case ST_AlgTypeTypeAny :_fee .Value ="\u0074y\u0070\u0065\u0041\u006e\u0079";case ST_AlgTypeCustom :_fee .Value ="\u0063\u0075\u0073\u0074\u006f\u006d";};return _fee ,nil ;};type ST_CryptProv byte ;func (_fgeb *ST_TrueFalseBlank )UnmarshalXML (d *_e .Decoder ,start _e .StartElement )error {_cce ,_dfaa :=d .Token ();
if _dfaa !=nil {return _dfaa ;};if _fdg ,_gdg :=_cce .(_e .EndElement );_gdg &&_fdg .Name ==start .Name {*_fgeb =1;return nil ;};if _ddg ,_cee :=_cce .(_e .CharData );!_cee {return _g .Errorf ("\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0063\u0068a\u0072\u0020\u0064\u0061\u0074\u0061\u002c\u0020\u0067\u006ft\u0020\u0025\u0054",_cce );
}else {switch string (_ddg ){case "":*_fgeb =0;case "\u0074":*_fgeb =1;case "\u0066":*_fgeb =2;case "\u0074\u0072\u0075\u0065":*_fgeb =3;case "\u0066\u0061\u006cs\u0065":*_fgeb =4;case "\u0054\u0072\u0075\u0065":*_fgeb =6;case "\u0046\u0061\u006cs\u0065":*_fgeb =7;
};};_cce ,_dfaa =d .Token ();if _dfaa !=nil {return _dfaa ;};if _cge ,_eeg :=_cce .(_e .EndElement );_eeg &&_cge .Name ==start .Name {return nil ;};return _g .Errorf ("\u0065\u0078\u0070\u0065c\u0074\u0065\u0064\u0020\u0065\u006e\u0064\u0020\u0065\u006ce\u006de\u006e\u0074\u002c\u0020\u0067\u006f\u0074 \u0025\u0076",_cce );
};type ST_ConformanceClass byte ;func (_bbb ST_VerticalAlignRun )MarshalXML (e *_e .Encoder ,start _e .StartElement )error {return e .EncodeElement (_bbb .String (),start );};const ST_PositiveUniversalMeasurePattern ="\u005b\u0030-9\u005d\u002b\u0028\\\u002e\u005b\u0030\u002d9]+\u0029?(\u006d\u006d\u007c\u0063\u006d\u007c\u0069n|\u0070\u0074\u007c\u0070\u0063\u007c\u0070i\u0029";
func (_gd ST_TwipsMeasure )String ()string {if _gd .ST_UnsignedDecimalNumber !=nil {return _g .Sprintf ("\u0025\u0076",*_gd .ST_UnsignedDecimalNumber );};if _gd .ST_PositiveUniversalMeasure !=nil {return _g .Sprintf ("\u0025\u0076",*_gd .ST_PositiveUniversalMeasure );
};return "";};func (_cba ST_TrueFalse )String ()string {switch _cba {case 0:return "";case 1:return "\u0074";case 2:return "\u0066";case 3:return "\u0074\u0072\u0075\u0065";case 4:return "\u0066\u0061\u006cs\u0065";};return "";};var ST_GuidPatternRe =_d .MustCompile (ST_GuidPattern );
func (_bcec ST_YAlign )ValidateWithPath (path string )error {switch _bcec {case 0,1,2,3,4,5,6:default:return _g .Errorf ("\u0025s\u003a\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072a\u006eg\u0065 \u0076\u0061\u006c\u0075\u0065\u0020\u0025d",path ,int (_bcec ));
};return nil ;};const (ST_CalendarTypeUnset ST_CalendarType =0;ST_CalendarTypeGregorian ST_CalendarType =1;ST_CalendarTypeGregorianUs ST_CalendarType =2;ST_CalendarTypeGregorianMeFrench ST_CalendarType =3;ST_CalendarTypeGregorianArabic ST_CalendarType =4;
ST_CalendarTypeHijri ST_CalendarType =5;ST_CalendarTypeHebrew ST_CalendarType =6;ST_CalendarTypeTaiwan ST_CalendarType =7;ST_CalendarTypeJapan ST_CalendarType =8;ST_CalendarTypeThai ST_CalendarType =9;ST_CalendarTypeKorea ST_CalendarType =10;ST_CalendarTypeSaka ST_CalendarType =11;
ST_CalendarTypeGregorianXlitEnglish ST_CalendarType =12;ST_CalendarTypeGregorianXlitFrench ST_CalendarType =13;ST_CalendarTypeNone ST_CalendarType =14;);func (_gg *ST_CalendarType )UnmarshalXML (d *_e .Decoder ,start _e .StartElement )error {_gb ,_fa :=d .Token ();
if _fa !=nil {return _fa ;};if _cg ,_cb :=_gb .(_e .EndElement );_cb &&_cg .Name ==start .Name {*_gg =1;return nil ;};if _ae ,_db :=_gb .(_e .CharData );!_db {return _g .Errorf ("\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0063\u0068a\u0072\u0020\u0064\u0061\u0074\u0061\u002c\u0020\u0067\u006ft\u0020\u0025\u0054",_gb );
}else {switch string (_ae ){case "":*_gg =0;case "\u0067r\u0065\u0067\u006f\u0072\u0069\u0061n":*_gg =1;case "g\u0072\u0065\u0067\u006f\u0072\u0069\u0061\u006e\u0055\u0073":*_gg =2;case "\u0067\u0072\u0065\u0067\u006f\u0072\u0069\u0061\u006e\u004d\u0065\u0046r\u0065\u006e\u0063\u0068":*_gg =3;
case "\u0067r\u0065g\u006f\u0072\u0069\u0061\u006e\u0041\u0072\u0061\u0062\u0069\u0063":*_gg =4;case "\u0068\u0069\u006ar\u0069":*_gg =5;case "\u0068\u0065\u0062\u0072\u0065\u0077":*_gg =6;case "\u0074\u0061\u0069\u0077\u0061\u006e":*_gg =7;case "\u006a\u0061\u0070a\u006e":*_gg =8;
case "\u0074\u0068\u0061\u0069":*_gg =9;case "\u006b\u006f\u0072e\u0061":*_gg =10;case "\u0073\u0061\u006b\u0061":*_gg =11;case "g\u0072e\u0067\u006f\u0072\u0069\u0061\u006e\u0058\u006ci\u0074\u0045\u006e\u0067li\u0073\u0068":*_gg =12;case "\u0067\u0072\u0065\u0067or\u0069\u0061\u006e\u0058\u006c\u0069\u0074\u0046\u0072\u0065\u006e\u0063\u0068":*_gg =13;
case "\u006e\u006f\u006e\u0065":*_gg =14;};};_gb ,_fa =d .Token ();if _fa !=nil {return _fa ;};if _be ,_fgb :=_gb .(_e .EndElement );_fgb &&_be .Name ==start .Name {return nil ;};return _g .Errorf ("\u0065\u0078\u0070\u0065c\u0074\u0065\u0064\u0020\u0065\u006e\u0064\u0020\u0065\u006ce\u006de\u006e\u0074\u002c\u0020\u0067\u006f\u0074 \u0025\u0076",_gb );
};func (_cfd ST_YAlign )MarshalXMLAttr (name _e .Name )(_e .Attr ,error ){_eff :=_e .Attr {};_eff .Name =name ;switch _cfd {case ST_YAlignUnset :_eff .Value ="";case ST_YAlignInline :_eff .Value ="\u0069\u006e\u006c\u0069\u006e\u0065";case ST_YAlignTop :_eff .Value ="\u0074\u006f\u0070";
case ST_YAlignCenter :_eff .Value ="\u0063\u0065\u006e\u0074\u0065\u0072";case ST_YAlignBottom :_eff .Value ="\u0062\u006f\u0074\u0074\u006f\u006d";case ST_YAlignInside :_eff .Value ="\u0069\u006e\u0073\u0069\u0064\u0065";case ST_YAlignOutside :_eff .Value ="\u006fu\u0074\u0073\u0069\u0064\u0065";
};return _eff ,nil ;};func (_fcf *ST_TrueFalse )UnmarshalXMLAttr (attr _e .Attr )error {switch attr .Value {case "":*_fcf =0;case "\u0074":*_fcf =1;case "\u0066":*_fcf =2;case "\u0074\u0072\u0075\u0065":*_fcf =3;case "\u0066\u0061\u006cs\u0065":*_fcf =4;
};return nil ;};func (_caf ST_AlgType )String ()string {switch _caf {case 0:return "";case 1:return "\u0074y\u0070\u0065\u0041\u006e\u0079";case 2:return "\u0063\u0075\u0073\u0074\u006f\u006d";};return "";};func (_aeb ST_ConformanceClass )MarshalXML (e *_e .Encoder ,start _e .StartElement )error {return e .EncodeElement (_aeb .String (),start );
};func (_gfg ST_CalendarType )Validate ()error {return _gfg .ValidateWithPath ("")};type ST_TrueFalseBlank byte ;func (_ccg ST_XAlign )MarshalXML (e *_e .Encoder ,start _e .StartElement )error {return e .EncodeElement (_ccg .String (),start );};func (_gcg ST_CalendarType )ValidateWithPath (path string )error {switch _gcg {case 0,1,2,3,4,5,6,7,8,9,10,11,12,13,14:default:return _g .Errorf ("\u0025s\u003a\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072a\u006eg\u0065 \u0076\u0061\u006c\u0075\u0065\u0020\u0025d",path ,int (_gcg ));
};return nil ;};func (_ce *ST_CryptProv )UnmarshalXML (d *_e .Decoder ,start _e .StartElement )error {_gbb ,_aef :=d .Token ();if _aef !=nil {return _aef ;};if _afa ,_cf :=_gbb .(_e .EndElement );_cf &&_afa .Name ==start .Name {*_ce =1;return nil ;};if _cdb ,_cdc :=_gbb .(_e .CharData );
!_cdc {return _g .Errorf ("\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0063\u0068a\u0072\u0020\u0064\u0061\u0074\u0061\u002c\u0020\u0067\u006ft\u0020\u0025\u0054",_gbb );}else {switch string (_cdb ){case "":*_ce =0;case "\u0072\u0073\u0061\u0041\u0045\u0053":*_ce =1;
case "\u0072s\u0061\u0046\u0075\u006c\u006c":*_ce =2;case "\u0063\u0075\u0073\u0074\u006f\u006d":*_ce =3;};};_gbb ,_aef =d .Token ();if _aef !=nil {return _aef ;};if _dfe ,_fd :=_gbb .(_e .EndElement );_fd &&_dfe .Name ==start .Name {return nil ;};return _g .Errorf ("\u0065\u0078\u0070\u0065c\u0074\u0065\u0064\u0020\u0065\u006e\u0064\u0020\u0065\u006ce\u006de\u006e\u0074\u002c\u0020\u0067\u006f\u0074 \u0025\u0076",_gbb );
};func (_acc ST_TrueFalseBlank )MarshalXML (e *_e .Encoder ,start _e .StartElement )error {return e .EncodeElement (_acc .String (),start );};var ST_UniversalMeasurePatternRe =_d .MustCompile (ST_UniversalMeasurePattern );func (_afaf ST_AlgType )Validate ()error {return _afaf .ValidateWithPath ("")};
func (_bce *ST_CalendarType )UnmarshalXMLAttr (attr _e .Attr )error {switch attr .Value {case "":*_bce =0;case "\u0067r\u0065\u0067\u006f\u0072\u0069\u0061n":*_bce =1;case "g\u0072\u0065\u0067\u006f\u0072\u0069\u0061\u006e\u0055\u0073":*_bce =2;case "\u0067\u0072\u0065\u0067\u006f\u0072\u0069\u0061\u006e\u004d\u0065\u0046r\u0065\u006e\u0063\u0068":*_bce =3;
case "\u0067r\u0065g\u006f\u0072\u0069\u0061\u006e\u0041\u0072\u0061\u0062\u0069\u0063":*_bce =4;case "\u0068\u0069\u006ar\u0069":*_bce =5;case "\u0068\u0065\u0062\u0072\u0065\u0077":*_bce =6;case "\u0074\u0061\u0069\u0077\u0061\u006e":*_bce =7;case "\u006a\u0061\u0070a\u006e":*_bce =8;
case "\u0074\u0068\u0061\u0069":*_bce =9;case "\u006b\u006f\u0072e\u0061":*_bce =10;case "\u0073\u0061\u006b\u0061":*_bce =11;case "g\u0072e\u0067\u006f\u0072\u0069\u0061\u006e\u0058\u006ci\u0074\u0045\u006e\u0067li\u0073\u0068":*_bce =12;case "\u0067\u0072\u0065\u0067or\u0069\u0061\u006e\u0058\u006c\u0069\u0074\u0046\u0072\u0065\u006e\u0063\u0068":*_bce =13;
case "\u006e\u006f\u006e\u0065":*_bce =14;};return nil ;};type ST_CalendarType byte ;func (_ge *ST_OnOff1 )UnmarshalXMLAttr (attr _e .Attr )error {switch attr .Value {case "":*_ge =0;case "\u006f\u006e":*_ge =1;case "\u006f\u0066\u0066":*_ge =2;};return nil ;
};func (_debb ST_VerticalAlignRun )Validate ()error {return _debb .ValidateWithPath ("")};func (_de *ST_OnOff )ValidateWithPath (path string )error {_ab :=[]string {};if _de .Bool !=nil {_ab =append (_ab ,"\u0042\u006f\u006f\u006c");};if _de .ST_OnOff1 !=ST_OnOff1Unset {_ab =append (_ab ,"\u0053T\u005f\u004f\u006e\u004f\u0066\u00661");
};if len (_ab )> 1{return _g .Errorf ("%\u0073\u0020\u0074\u006f\u006f\u0020m\u0061\u006e\u0079\u0020\u006d\u0065\u006d\u0062\u0065r\u0073\u0020\u0073e\u0074:\u0020\u0025\u0076",path ,_ab );};return nil ;};const ST_FixedPercentagePattern ="-\u003f\u0028\u0028\u0031\u0030\u0030\u0029\u007c\u0028\u005b\u0030\u002d\u0039\u005d\u005b\u0030\u002d\u0039]\u003f\u0029\u0029\u0028\u005c\u002e\u005b\u0030\u002d\u0039][\u0030\u002d\u0039]\u003f)\u003f\u0025";
func (_abf ST_CryptProv )MarshalXMLAttr (name _e .Name )(_e .Attr ,error ){_bcg :=_e .Attr {};_bcg .Name =name ;switch _abf {case ST_CryptProvUnset :_bcg .Value ="";case ST_CryptProvRsaAES :_bcg .Value ="\u0072\u0073\u0061\u0041\u0045\u0053";case ST_CryptProvRsaFull :_bcg .Value ="\u0072s\u0061\u0046\u0075\u006c\u006c";
case ST_CryptProvCustom :_bcg .Value ="\u0063\u0075\u0073\u0074\u006f\u006d";};return _bcg ,nil ;};const (ST_TrueFalseBlankUnset ST_TrueFalseBlank =0;ST_TrueFalseBlankT ST_TrueFalseBlank =1;ST_TrueFalseBlankF ST_TrueFalseBlank =2;ST_TrueFalseBlankTrue ST_TrueFalseBlank =3;
ST_TrueFalseBlankFalse ST_TrueFalseBlank =4;ST_TrueFalseBlankTrue_ ST_TrueFalseBlank =6;ST_TrueFalseBlankFalse_ ST_TrueFalseBlank =7;);func (_ea ST_TrueFalseBlank )MarshalXMLAttr (name _e .Name )(_e .Attr ,error ){_ddee :=_e .Attr {};_ddee .Name =name ;
switch _ea {case ST_TrueFalseBlankUnset :_ddee .Value ="";case ST_TrueFalseBlankT :_ddee .Value ="\u0074";case ST_TrueFalseBlankF :_ddee .Value ="\u0066";case ST_TrueFalseBlankTrue :_ddee .Value ="\u0074\u0072\u0075\u0065";case ST_TrueFalseBlankFalse :_ddee .Value ="\u0066\u0061\u006cs\u0065";
case ST_TrueFalseBlankTrue_ :_ddee .Value ="\u0054\u0072\u0075\u0065";case ST_TrueFalseBlankFalse_ :_ddee .Value ="\u0046\u0061\u006cs\u0065";};return _ddee ,nil ;};func (_deb ST_TrueFalse )MarshalXMLAttr (name _e .Name )(_e .Attr ,error ){_fec :=_e .Attr {};
_fec .Name =name ;switch _deb {case ST_TrueFalseUnset :_fec .Value ="";case ST_TrueFalseT :_fec .Value ="\u0074";case ST_TrueFalseF :_fec .Value ="\u0066";case ST_TrueFalseTrue :_fec .Value ="\u0074\u0072\u0075\u0065";case ST_TrueFalseFalse :_fec .Value ="\u0066\u0061\u006cs\u0065";
};return _fec ,nil ;};func (_beb *ST_CryptProv )UnmarshalXMLAttr (attr _e .Attr )error {switch attr .Value {case "":*_beb =0;case "\u0072\u0073\u0061\u0041\u0045\u0053":*_beb =1;case "\u0072s\u0061\u0046\u0075\u006c\u006c":*_beb =2;case "\u0063\u0075\u0073\u0074\u006f\u006d":*_beb =3;
};return nil ;};const (ST_AlgTypeUnset ST_AlgType =0;ST_AlgTypeTypeAny ST_AlgType =1;ST_AlgTypeCustom ST_AlgType =2;);func (_aac ST_OnOff1 )Validate ()error {return _aac .ValidateWithPath ("")};type ST_OnOff1 byte ;func (_dgg ST_TrueFalse )MarshalXML (e *_e .Encoder ,start _e .StartElement )error {return e .EncodeElement (_dgg .String (),start );
};func (_bc ST_OnOff )String ()string {if _bc .Bool !=nil {return _g .Sprintf ("\u0025\u0076",*_bc .Bool );};if _bc .ST_OnOff1 !=ST_OnOff1Unset {return _bc .ST_OnOff1 .String ();};return "";};func (_gdag ST_ConformanceClass )Validate ()error {return _gdag .ValidateWithPath ("")};
func (_cca ST_YAlign )MarshalXML (e *_e .Encoder ,start _e .StartElement )error {return e .EncodeElement (_cca .String (),start );};func (_dfa ST_CryptProv )Validate ()error {return _dfa .ValidateWithPath ("")};func (_ff *ST_OnOff1 )UnmarshalXML (d *_e .Decoder ,start _e .StartElement )error {_bff ,_cfg :=d .Token ();
if _cfg !=nil {return _cfg ;};if _eeb ,_cab :=_bff .(_e .EndElement );_cab &&_eeb .Name ==start .Name {*_ff =1;return nil ;};if _gbfe ,_gge :=_bff .(_e .CharData );!_gge {return _g .Errorf ("\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0063\u0068a\u0072\u0020\u0064\u0061\u0074\u0061\u002c\u0020\u0067\u006ft\u0020\u0025\u0054",_bff );
}else {switch string (_gbfe ){case "":*_ff =0;case "\u006f\u006e":*_ff =1;case "\u006f\u0066\u0066":*_ff =2;};};_bff ,_cfg =d .Token ();if _cfg !=nil {return _cfg ;};if _ffa ,_bfa :=_bff .(_e .EndElement );_bfa &&_ffa .Name ==start .Name {return nil ;};
return _g .Errorf ("\u0065\u0078\u0070\u0065c\u0074\u0065\u0064\u0020\u0065\u006e\u0064\u0020\u0065\u006ce\u006de\u006e\u0074\u002c\u0020\u0067\u006f\u0074 \u0025\u0076",_bff );};var ST_PositiveFixedPercentagePatternRe =_d .MustCompile (ST_PositiveFixedPercentagePattern );
func (_caa ST_ConformanceClass )MarshalXMLAttr (name _e .Name )(_e .Attr ,error ){_edc :=_e .Attr {};_edc .Name =name ;switch _caa {case ST_ConformanceClassUnset :_edc .Value ="";case ST_ConformanceClassStrict :_edc .Value ="\u0073\u0074\u0072\u0069\u0063\u0074";
case ST_ConformanceClassTransitional :_edc .Value ="\u0074\u0072\u0061n\u0073\u0069\u0074\u0069\u006f\u006e\u0061\u006c";};return _edc ,nil ;};func (_gffd ST_VerticalAlignRun )ValidateWithPath (path string )error {switch _gffd {case 0,1,2,3:default:return _g .Errorf ("\u0025s\u003a\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072a\u006eg\u0065 \u0076\u0061\u006c\u0075\u0065\u0020\u0025d",path ,int (_gffd ));
};return nil ;};func _gdd (_da bool )uint8 {if _da {return 1;};return 0;};func (_ffc *ST_XAlign )UnmarshalXML (d *_e .Decoder ,start _e .StartElement )error {_dag ,_cea :=d .Token ();if _cea !=nil {return _cea ;};if _ggeg ,_eac :=_dag .(_e .EndElement );
_eac &&_ggeg .Name ==start .Name {*_ffc =1;return nil ;};if _aga ,_gea :=_dag .(_e .CharData );!_gea {return _g .Errorf ("\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0063\u0068a\u0072\u0020\u0064\u0061\u0074\u0061\u002c\u0020\u0067\u006ft\u0020\u0025\u0054",_dag );
}else {switch string (_aga ){case "":*_ffc =0;case "\u006c\u0065\u0066\u0074":*_ffc =1;case "\u0063\u0065\u006e\u0074\u0065\u0072":*_ffc =2;case "\u0072\u0069\u0067h\u0074":*_ffc =3;case "\u0069\u006e\u0073\u0069\u0064\u0065":*_ffc =4;case "\u006fu\u0074\u0073\u0069\u0064\u0065":*_ffc =5;
};};_dag ,_cea =d .Token ();if _cea !=nil {return _cea ;};if _bdaa ,_fgef :=_dag .(_e .EndElement );_fgef &&_bdaa .Name ==start .Name {return nil ;};return _g .Errorf ("\u0065\u0078\u0070\u0065c\u0074\u0065\u0064\u0020\u0065\u006e\u0064\u0020\u0065\u006ce\u006de\u006e\u0074\u002c\u0020\u0067\u006f\u0074 \u0025\u0076",_dag );
};func (_dde ST_CryptProv )MarshalXML (e *_e .Encoder ,start _e .StartElement )error {return e .EncodeElement (_dde .String (),start );};func (_ee ST_OnOff1 )MarshalXMLAttr (name _e .Name )(_e .Attr ,error ){_dg :=_e .Attr {};_dg .Name =name ;switch _ee {case ST_OnOff1Unset :_dg .Value ="";
case ST_OnOff1On :_dg .Value ="\u006f\u006e";case ST_OnOff1Off :_dg .Value ="\u006f\u0066\u0066";};return _dg ,nil ;};func (_gde *ST_AlgClass )UnmarshalXMLAttr (attr _e .Attr )error {switch attr .Value {case "":*_gde =0;case "\u0068\u0061\u0073\u0068":*_gde =1;
case "\u0063\u0075\u0073\u0074\u006f\u006d":*_gde =2;};return nil ;};func (_fed ST_ConformanceClass )String ()string {switch _fed {case 0:return "";case 1:return "\u0073\u0074\u0072\u0069\u0063\u0074";case 2:return "\u0074\u0072\u0061n\u0073\u0069\u0074\u0069\u006f\u006e\u0061\u006c";
};return "";};func (_gae *ST_YAlign )UnmarshalXML (d *_e .Decoder ,start _e .StartElement )error {_gdb ,_geg :=d .Token ();if _geg !=nil {return _geg ;};if _ggc ,_dac :=_gdb .(_e .EndElement );_dac &&_ggc .Name ==start .Name {*_gae =1;return nil ;};if _agc ,_ecg :=_gdb .(_e .CharData );
!_ecg {return _g .Errorf ("\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0063\u0068a\u0072\u0020\u0064\u0061\u0074\u0061\u002c\u0020\u0067\u006ft\u0020\u0025\u0054",_gdb );}else {switch string (_agc ){case "":*_gae =0;case "\u0069\u006e\u006c\u0069\u006e\u0065":*_gae =1;
case "\u0074\u006f\u0070":*_gae =2;case "\u0063\u0065\u006e\u0074\u0065\u0072":*_gae =3;case "\u0062\u006f\u0074\u0074\u006f\u006d":*_gae =4;case "\u0069\u006e\u0073\u0069\u0064\u0065":*_gae =5;case "\u006fu\u0074\u0073\u0069\u0064\u0065":*_gae =6;};};
_gdb ,_geg =d .Token ();if _geg !=nil {return _geg ;};if _ceee ,_efe :=_gdb .(_e .EndElement );_efe &&_ceee .Name ==start .Name {return nil ;};return _g .Errorf ("\u0065\u0078\u0070\u0065c\u0074\u0065\u0064\u0020\u0065\u006e\u0064\u0020\u0065\u006ce\u006de\u006e\u0074\u002c\u0020\u0067\u006f\u0074 \u0025\u0076",_gdb );
};const (ST_ConformanceClassUnset ST_ConformanceClass =0;ST_ConformanceClassStrict ST_ConformanceClass =1;ST_ConformanceClassTransitional ST_ConformanceClass =2;);func (_eg ST_CalendarType )MarshalXML (e *_e .Encoder ,start _e .StartElement )error {return e .EncodeElement (_eg .String (),start );
};var ST_PercentagePatternRe =_d .MustCompile (ST_PercentagePattern );func (_ef ST_CryptProv )String ()string {switch _ef {case 0:return "";case 1:return "\u0072\u0073\u0061\u0041\u0045\u0053";case 2:return "\u0072s\u0061\u0046\u0075\u006c\u006c";case 3:return "\u0063\u0075\u0073\u0074\u006f\u006d";
};return "";};func (_ega *ST_XAlign )UnmarshalXMLAttr (attr _e .Attr )error {switch attr .Value {case "":*_ega =0;case "\u006c\u0065\u0066\u0074":*_ega =1;case "\u0063\u0065\u006e\u0074\u0065\u0072":*_ega =2;case "\u0072\u0069\u0067h\u0074":*_ega =3;case "\u0069\u006e\u0073\u0069\u0064\u0065":*_ega =4;
case "\u006fu\u0074\u0073\u0069\u0064\u0065":*_ega =5;};return nil ;};func (_fc ST_AlgClass )ValidateWithPath (path string )error {switch _fc {case 0,1,2:default:return _g .Errorf ("\u0025s\u003a\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072a\u006eg\u0065 \u0076\u0061\u006c\u0075\u0065\u0020\u0025d",path ,int (_fc ));
};return nil ;};func (_gff *ST_AlgType )UnmarshalXML (d *_e .Decoder ,start _e .StartElement )error {_aa ,_agf :=d .Token ();if _agf !=nil {return _agf ;};if _fef ,_egce :=_aa .(_e .EndElement );_egce &&_fef .Name ==start .Name {*_gff =1;return nil ;};
if _fca ,_fea :=_aa .(_e .CharData );!_fea {return _g .Errorf ("\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0063\u0068a\u0072\u0020\u0064\u0061\u0074\u0061\u002c\u0020\u0067\u006ft\u0020\u0025\u0054",_aa );}else {switch string (_fca ){case "":*_gff =0;
case "\u0074y\u0070\u0065\u0041\u006e\u0079":*_gff =1;case "\u0063\u0075\u0073\u0074\u006f\u006d":*_gff =2;};};_aa ,_agf =d .Token ();if _agf !=nil {return _agf ;};if _gbe ,_acg :=_aa .(_e .EndElement );_acg &&_gbe .Name ==start .Name {return nil ;};return _g .Errorf ("\u0065\u0078\u0070\u0065c\u0074\u0065\u0064\u0020\u0065\u006e\u0064\u0020\u0065\u006ce\u006de\u006e\u0074\u002c\u0020\u0067\u006f\u0074 \u0025\u0076",_aa );
};const ST_GuidPattern ="\u005c\u007b\u005b\u0030\u002d\u0039\u0041\u002d\u0046\u005d\u007b\u0038\u007d\u002d\u005b\u0030\u002d9\u0041\u002d\u0046\u005d\u007b\u0034\u007d\u002d\u005b\u0030-\u0039\u0041\u002d\u0046\u005d\u007b\u0034\u007d\u002d\u005b\u0030\u002d\u0039\u0041\u002d\u0046\u005d\u007b4\u007d\u002d\u005b\u0030\u002d\u0039A\u002d\u0046]\u007b\u00312\u007d\\\u007d";
const ST_PositiveFixedPercentagePattern ="\u0028\u0028\u0031\u0030\u0030\u0029\u007c\u0028\u005b\u0030\u002d\u0039\u005d\u005b\u0030\u002d\u0039\u005d\u003f\u0029\u0029\u0028\u005c\u002e[\u0030\u002d\u0039\u005d\u005b0\u002d\u0039]\u003f\u0029\u003f\u0025";