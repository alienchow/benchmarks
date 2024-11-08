package main

import (
	"errors"
	"fmt"
	"math"
	"sort"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

var (
	percentiles = map[string][]int64{}

	configs = map[string]bool{
		"feature_0": true,
		"feature_1": false,
	}

	largeConfigs = map[string]bool{
		"feature_0":    true,
		"feature_1":    false,
		"feature_2":    false,
		"feature_3":    false,
		"feature_4":    false,
		"feature_5":    false,
		"feature_6":    false,
		"feature_7":    false,
		"feature_8":    false,
		"feature_9":    false,
		"feature_10":   false,
		"feature_11":   false,
		"feature_12":   false,
		"feature_13":   false,
		"feature_14":   false,
		"feature_15":   false,
		"feature_16":   false,
		"feature_17":   false,
		"feature_18":   false,
		"feature_19":   false,
		"feature_20":   false,
		"feature_21":   false,
		"feature_22":   false,
		"feature_23":   false,
		"feature_24":   false,
		"feature_25":   false,
		"feature_26":   false,
		"feature_27":   false,
		"feature_28":   false,
		"feature_29":   false,
		"feature_30":   false,
		"feature_31":   false,
		"feature_32":   false,
		"feature_33":   false,
		"feature_34":   false,
		"feature_35":   false,
		"feature_36":   false,
		"feature_37":   false,
		"feature_38":   false,
		"feature_39":   false,
		"feature_40":   false,
		"feature_41":   false,
		"feature_42":   false,
		"feature_43":   false,
		"feature_44":   false,
		"feature_45":   false,
		"feature_46":   false,
		"feature_47":   false,
		"feature_48":   false,
		"feature_49":   false,
		"feature_50":   false,
		"feature_51":   false,
		"feature_52":   false,
		"feature_53":   false,
		"feature_54":   false,
		"feature_55":   false,
		"feature_56":   false,
		"feature_57":   false,
		"feature_58":   false,
		"feature_59":   false,
		"feature_60":   false,
		"feature_61":   false,
		"feature_62":   false,
		"feature_63":   false,
		"feature_64":   false,
		"feature_65":   false,
		"feature_66":   false,
		"feature_67":   false,
		"feature_68":   false,
		"feature_69":   false,
		"feature_70":   false,
		"feature_71":   false,
		"feature_72":   false,
		"feature_73":   false,
		"feature_74":   false,
		"feature_75":   false,
		"feature_76":   false,
		"feature_77":   false,
		"feature_78":   false,
		"feature_79":   false,
		"feature_80":   false,
		"feature_81":   false,
		"feature_82":   false,
		"feature_83":   false,
		"feature_84":   false,
		"feature_85":   false,
		"feature_86":   false,
		"feature_87":   false,
		"feature_88":   false,
		"feature_89":   false,
		"feature_90":   false,
		"feature_91":   false,
		"feature_92":   false,
		"feature_93":   false,
		"feature_94":   false,
		"feature_95":   false,
		"feature_96":   false,
		"feature_97":   false,
		"feature_98":   false,
		"feature_99":   false,
		"feature_100":  false,
		"feature_101":  false,
		"feature_102":  false,
		"feature_103":  false,
		"feature_104":  false,
		"feature_105":  false,
		"feature_106":  false,
		"feature_107":  false,
		"feature_108":  false,
		"feature_109":  false,
		"feature_110":  false,
		"feature_111":  false,
		"feature_112":  false,
		"feature_113":  false,
		"feature_114":  false,
		"feature_115":  false,
		"feature_116":  false,
		"feature_117":  false,
		"feature_118":  false,
		"feature_119":  false,
		"feature_120":  false,
		"feature_121":  false,
		"feature_122":  false,
		"feature_123":  false,
		"feature_124":  false,
		"feature_125":  false,
		"feature_126":  false,
		"feature_127":  false,
		"feature_128":  false,
		"feature_129":  false,
		"feature_130":  false,
		"feature_131":  false,
		"feature_132":  false,
		"feature_133":  false,
		"feature_134":  false,
		"feature_135":  false,
		"feature_136":  false,
		"feature_137":  false,
		"feature_138":  false,
		"feature_139":  false,
		"feature_140":  false,
		"feature_141":  false,
		"feature_142":  false,
		"feature_143":  false,
		"feature_144":  false,
		"feature_145":  false,
		"feature_146":  false,
		"feature_147":  false,
		"feature_148":  false,
		"feature_149":  false,
		"feature_150":  false,
		"feature_151":  false,
		"feature_152":  false,
		"feature_153":  false,
		"feature_154":  false,
		"feature_155":  false,
		"feature_156":  false,
		"feature_157":  false,
		"feature_158":  false,
		"feature_159":  false,
		"feature_160":  false,
		"feature_161":  false,
		"feature_162":  false,
		"feature_163":  false,
		"feature_164":  false,
		"feature_165":  false,
		"feature_166":  false,
		"feature_167":  false,
		"feature_168":  false,
		"feature_169":  false,
		"feature_170":  false,
		"feature_171":  false,
		"feature_172":  false,
		"feature_173":  false,
		"feature_174":  false,
		"feature_175":  false,
		"feature_176":  false,
		"feature_177":  false,
		"feature_178":  false,
		"feature_179":  false,
		"feature_180":  false,
		"feature_181":  false,
		"feature_182":  false,
		"feature_183":  false,
		"feature_184":  false,
		"feature_185":  false,
		"feature_186":  false,
		"feature_187":  false,
		"feature_188":  false,
		"feature_189":  false,
		"feature_190":  false,
		"feature_191":  false,
		"feature_192":  false,
		"feature_193":  false,
		"feature_194":  false,
		"feature_195":  false,
		"feature_196":  false,
		"feature_197":  false,
		"feature_198":  false,
		"feature_199":  false,
		"feature_200":  false,
		"feature_201":  false,
		"feature_202":  false,
		"feature_203":  false,
		"feature_204":  false,
		"feature_205":  false,
		"feature_206":  false,
		"feature_207":  false,
		"feature_208":  false,
		"feature_209":  false,
		"feature_210":  false,
		"feature_211":  false,
		"feature_212":  false,
		"feature_213":  false,
		"feature_214":  false,
		"feature_215":  false,
		"feature_216":  false,
		"feature_217":  false,
		"feature_218":  false,
		"feature_219":  false,
		"feature_220":  false,
		"feature_221":  false,
		"feature_222":  false,
		"feature_223":  false,
		"feature_224":  false,
		"feature_225":  false,
		"feature_226":  false,
		"feature_227":  false,
		"feature_228":  false,
		"feature_229":  false,
		"feature_230":  false,
		"feature_231":  false,
		"feature_232":  false,
		"feature_233":  false,
		"feature_234":  false,
		"feature_235":  false,
		"feature_236":  false,
		"feature_237":  false,
		"feature_238":  false,
		"feature_239":  false,
		"feature_240":  false,
		"feature_241":  false,
		"feature_242":  false,
		"feature_243":  false,
		"feature_244":  false,
		"feature_245":  false,
		"feature_246":  false,
		"feature_247":  false,
		"feature_248":  false,
		"feature_249":  false,
		"feature_250":  false,
		"feature_251":  false,
		"feature_252":  false,
		"feature_253":  false,
		"feature_254":  false,
		"feature_255":  false,
		"feature_256":  false,
		"feature_257":  false,
		"feature_258":  false,
		"feature_259":  false,
		"feature_260":  false,
		"feature_261":  false,
		"feature_262":  false,
		"feature_263":  false,
		"feature_264":  false,
		"feature_265":  false,
		"feature_266":  false,
		"feature_267":  false,
		"feature_268":  false,
		"feature_269":  false,
		"feature_270":  false,
		"feature_271":  false,
		"feature_272":  false,
		"feature_273":  false,
		"feature_274":  false,
		"feature_275":  false,
		"feature_276":  false,
		"feature_277":  false,
		"feature_278":  false,
		"feature_279":  false,
		"feature_280":  false,
		"feature_281":  false,
		"feature_282":  false,
		"feature_283":  false,
		"feature_284":  false,
		"feature_285":  false,
		"feature_286":  false,
		"feature_287":  false,
		"feature_288":  false,
		"feature_289":  false,
		"feature_290":  false,
		"feature_291":  false,
		"feature_292":  false,
		"feature_293":  false,
		"feature_294":  false,
		"feature_295":  false,
		"feature_296":  false,
		"feature_297":  false,
		"feature_298":  false,
		"feature_299":  false,
		"feature_300":  false,
		"feature_301":  false,
		"feature_302":  false,
		"feature_303":  false,
		"feature_304":  false,
		"feature_305":  false,
		"feature_306":  false,
		"feature_307":  false,
		"feature_308":  false,
		"feature_309":  false,
		"feature_310":  false,
		"feature_311":  false,
		"feature_312":  false,
		"feature_313":  false,
		"feature_314":  false,
		"feature_315":  false,
		"feature_316":  false,
		"feature_317":  false,
		"feature_318":  false,
		"feature_319":  false,
		"feature_320":  false,
		"feature_321":  false,
		"feature_322":  false,
		"feature_323":  false,
		"feature_324":  false,
		"feature_325":  false,
		"feature_326":  false,
		"feature_327":  false,
		"feature_328":  false,
		"feature_329":  false,
		"feature_330":  false,
		"feature_331":  false,
		"feature_332":  false,
		"feature_333":  false,
		"feature_334":  false,
		"feature_335":  false,
		"feature_336":  false,
		"feature_337":  false,
		"feature_338":  false,
		"feature_339":  false,
		"feature_340":  false,
		"feature_341":  false,
		"feature_342":  false,
		"feature_343":  false,
		"feature_344":  false,
		"feature_345":  false,
		"feature_346":  false,
		"feature_347":  false,
		"feature_348":  false,
		"feature_349":  false,
		"feature_350":  false,
		"feature_351":  false,
		"feature_352":  false,
		"feature_353":  false,
		"feature_354":  false,
		"feature_355":  false,
		"feature_356":  false,
		"feature_357":  false,
		"feature_358":  false,
		"feature_359":  false,
		"feature_360":  false,
		"feature_361":  false,
		"feature_362":  false,
		"feature_363":  false,
		"feature_364":  false,
		"feature_365":  false,
		"feature_366":  false,
		"feature_367":  false,
		"feature_368":  false,
		"feature_369":  false,
		"feature_370":  false,
		"feature_371":  false,
		"feature_372":  false,
		"feature_373":  false,
		"feature_374":  false,
		"feature_375":  false,
		"feature_376":  false,
		"feature_377":  false,
		"feature_378":  false,
		"feature_379":  false,
		"feature_380":  false,
		"feature_381":  false,
		"feature_382":  false,
		"feature_383":  false,
		"feature_384":  false,
		"feature_385":  false,
		"feature_386":  false,
		"feature_387":  false,
		"feature_388":  false,
		"feature_389":  false,
		"feature_390":  false,
		"feature_391":  false,
		"feature_392":  false,
		"feature_393":  false,
		"feature_394":  false,
		"feature_395":  false,
		"feature_396":  false,
		"feature_397":  false,
		"feature_398":  false,
		"feature_399":  false,
		"feature_400":  false,
		"feature_401":  false,
		"feature_402":  false,
		"feature_403":  false,
		"feature_404":  false,
		"feature_405":  false,
		"feature_406":  false,
		"feature_407":  false,
		"feature_408":  false,
		"feature_409":  false,
		"feature_410":  false,
		"feature_411":  false,
		"feature_412":  false,
		"feature_413":  false,
		"feature_414":  false,
		"feature_415":  false,
		"feature_416":  false,
		"feature_417":  false,
		"feature_418":  false,
		"feature_419":  false,
		"feature_420":  false,
		"feature_421":  false,
		"feature_422":  false,
		"feature_423":  false,
		"feature_424":  false,
		"feature_425":  false,
		"feature_426":  false,
		"feature_427":  false,
		"feature_428":  false,
		"feature_429":  false,
		"feature_430":  false,
		"feature_431":  false,
		"feature_432":  false,
		"feature_433":  false,
		"feature_434":  false,
		"feature_435":  false,
		"feature_436":  false,
		"feature_437":  false,
		"feature_438":  false,
		"feature_439":  false,
		"feature_440":  false,
		"feature_441":  false,
		"feature_442":  false,
		"feature_443":  false,
		"feature_444":  false,
		"feature_445":  false,
		"feature_446":  false,
		"feature_447":  false,
		"feature_448":  false,
		"feature_449":  false,
		"feature_450":  false,
		"feature_451":  false,
		"feature_452":  false,
		"feature_453":  false,
		"feature_454":  false,
		"feature_455":  false,
		"feature_456":  false,
		"feature_457":  false,
		"feature_458":  false,
		"feature_459":  false,
		"feature_460":  false,
		"feature_461":  false,
		"feature_462":  false,
		"feature_463":  false,
		"feature_464":  false,
		"feature_465":  false,
		"feature_466":  false,
		"feature_467":  false,
		"feature_468":  false,
		"feature_469":  false,
		"feature_470":  false,
		"feature_471":  false,
		"feature_472":  false,
		"feature_473":  false,
		"feature_474":  false,
		"feature_475":  false,
		"feature_476":  false,
		"feature_477":  false,
		"feature_478":  false,
		"feature_479":  false,
		"feature_480":  false,
		"feature_481":  false,
		"feature_482":  false,
		"feature_483":  false,
		"feature_484":  false,
		"feature_485":  false,
		"feature_486":  false,
		"feature_487":  false,
		"feature_488":  false,
		"feature_489":  false,
		"feature_490":  false,
		"feature_491":  false,
		"feature_492":  false,
		"feature_493":  false,
		"feature_494":  false,
		"feature_495":  false,
		"feature_496":  false,
		"feature_497":  false,
		"feature_498":  false,
		"feature_499":  false,
		"feature_500":  false,
		"feature_501":  false,
		"feature_502":  false,
		"feature_503":  false,
		"feature_504":  false,
		"feature_505":  false,
		"feature_506":  false,
		"feature_507":  false,
		"feature_508":  false,
		"feature_509":  false,
		"feature_510":  false,
		"feature_511":  false,
		"feature_512":  false,
		"feature_513":  false,
		"feature_514":  false,
		"feature_515":  false,
		"feature_516":  false,
		"feature_517":  false,
		"feature_518":  false,
		"feature_519":  false,
		"feature_520":  false,
		"feature_521":  false,
		"feature_522":  false,
		"feature_523":  false,
		"feature_524":  false,
		"feature_525":  false,
		"feature_526":  false,
		"feature_527":  false,
		"feature_528":  false,
		"feature_529":  false,
		"feature_530":  false,
		"feature_531":  false,
		"feature_532":  false,
		"feature_533":  false,
		"feature_534":  false,
		"feature_535":  false,
		"feature_536":  false,
		"feature_537":  false,
		"feature_538":  false,
		"feature_539":  false,
		"feature_540":  false,
		"feature_541":  false,
		"feature_542":  false,
		"feature_543":  false,
		"feature_544":  false,
		"feature_545":  false,
		"feature_546":  false,
		"feature_547":  false,
		"feature_548":  false,
		"feature_549":  false,
		"feature_550":  false,
		"feature_551":  false,
		"feature_552":  false,
		"feature_553":  false,
		"feature_554":  false,
		"feature_555":  false,
		"feature_556":  false,
		"feature_557":  false,
		"feature_558":  false,
		"feature_559":  false,
		"feature_560":  false,
		"feature_561":  false,
		"feature_562":  false,
		"feature_563":  false,
		"feature_564":  false,
		"feature_565":  false,
		"feature_566":  false,
		"feature_567":  false,
		"feature_568":  false,
		"feature_569":  false,
		"feature_570":  false,
		"feature_571":  false,
		"feature_572":  false,
		"feature_573":  false,
		"feature_574":  false,
		"feature_575":  false,
		"feature_576":  false,
		"feature_577":  false,
		"feature_578":  false,
		"feature_579":  false,
		"feature_580":  false,
		"feature_581":  false,
		"feature_582":  false,
		"feature_583":  false,
		"feature_584":  false,
		"feature_585":  false,
		"feature_586":  false,
		"feature_587":  false,
		"feature_588":  false,
		"feature_589":  false,
		"feature_590":  false,
		"feature_591":  false,
		"feature_592":  false,
		"feature_593":  false,
		"feature_594":  false,
		"feature_595":  false,
		"feature_596":  false,
		"feature_597":  false,
		"feature_598":  false,
		"feature_599":  false,
		"feature_600":  false,
		"feature_601":  false,
		"feature_602":  false,
		"feature_603":  false,
		"feature_604":  false,
		"feature_605":  false,
		"feature_606":  false,
		"feature_607":  false,
		"feature_608":  false,
		"feature_609":  false,
		"feature_610":  false,
		"feature_611":  false,
		"feature_612":  false,
		"feature_613":  false,
		"feature_614":  false,
		"feature_615":  false,
		"feature_616":  false,
		"feature_617":  false,
		"feature_618":  false,
		"feature_619":  false,
		"feature_620":  false,
		"feature_621":  false,
		"feature_622":  false,
		"feature_623":  false,
		"feature_624":  false,
		"feature_625":  false,
		"feature_626":  false,
		"feature_627":  false,
		"feature_628":  false,
		"feature_629":  false,
		"feature_630":  false,
		"feature_631":  false,
		"feature_632":  false,
		"feature_633":  false,
		"feature_634":  false,
		"feature_635":  false,
		"feature_636":  false,
		"feature_637":  false,
		"feature_638":  false,
		"feature_639":  false,
		"feature_640":  false,
		"feature_641":  false,
		"feature_642":  false,
		"feature_643":  false,
		"feature_644":  false,
		"feature_645":  false,
		"feature_646":  false,
		"feature_647":  false,
		"feature_648":  false,
		"feature_649":  false,
		"feature_650":  false,
		"feature_651":  false,
		"feature_652":  false,
		"feature_653":  false,
		"feature_654":  false,
		"feature_655":  false,
		"feature_656":  false,
		"feature_657":  false,
		"feature_658":  false,
		"feature_659":  false,
		"feature_660":  false,
		"feature_661":  false,
		"feature_662":  false,
		"feature_663":  false,
		"feature_664":  false,
		"feature_665":  false,
		"feature_666":  false,
		"feature_667":  false,
		"feature_668":  false,
		"feature_669":  false,
		"feature_670":  false,
		"feature_671":  false,
		"feature_672":  false,
		"feature_673":  false,
		"feature_674":  false,
		"feature_675":  false,
		"feature_676":  false,
		"feature_677":  false,
		"feature_678":  false,
		"feature_679":  false,
		"feature_680":  false,
		"feature_681":  false,
		"feature_682":  false,
		"feature_683":  false,
		"feature_684":  false,
		"feature_685":  false,
		"feature_686":  false,
		"feature_687":  false,
		"feature_688":  false,
		"feature_689":  false,
		"feature_690":  false,
		"feature_691":  false,
		"feature_692":  false,
		"feature_693":  false,
		"feature_694":  false,
		"feature_695":  false,
		"feature_696":  false,
		"feature_697":  false,
		"feature_698":  false,
		"feature_699":  false,
		"feature_700":  false,
		"feature_701":  false,
		"feature_702":  false,
		"feature_703":  false,
		"feature_704":  false,
		"feature_705":  false,
		"feature_706":  false,
		"feature_707":  false,
		"feature_708":  false,
		"feature_709":  false,
		"feature_710":  false,
		"feature_711":  false,
		"feature_712":  false,
		"feature_713":  false,
		"feature_714":  false,
		"feature_715":  false,
		"feature_716":  false,
		"feature_717":  false,
		"feature_718":  false,
		"feature_719":  false,
		"feature_720":  false,
		"feature_721":  false,
		"feature_722":  false,
		"feature_723":  false,
		"feature_724":  false,
		"feature_725":  false,
		"feature_726":  false,
		"feature_727":  false,
		"feature_728":  false,
		"feature_729":  false,
		"feature_730":  false,
		"feature_731":  false,
		"feature_732":  false,
		"feature_733":  false,
		"feature_734":  false,
		"feature_735":  false,
		"feature_736":  false,
		"feature_737":  false,
		"feature_738":  false,
		"feature_739":  false,
		"feature_740":  false,
		"feature_741":  false,
		"feature_742":  false,
		"feature_743":  false,
		"feature_744":  false,
		"feature_745":  false,
		"feature_746":  false,
		"feature_747":  false,
		"feature_748":  false,
		"feature_749":  false,
		"feature_750":  false,
		"feature_751":  false,
		"feature_752":  false,
		"feature_753":  false,
		"feature_754":  false,
		"feature_755":  false,
		"feature_756":  false,
		"feature_757":  false,
		"feature_758":  false,
		"feature_759":  false,
		"feature_760":  false,
		"feature_761":  false,
		"feature_762":  false,
		"feature_763":  false,
		"feature_764":  false,
		"feature_765":  false,
		"feature_766":  false,
		"feature_767":  false,
		"feature_768":  false,
		"feature_769":  false,
		"feature_770":  false,
		"feature_771":  false,
		"feature_772":  false,
		"feature_773":  false,
		"feature_774":  false,
		"feature_775":  false,
		"feature_776":  false,
		"feature_777":  false,
		"feature_778":  false,
		"feature_779":  false,
		"feature_780":  false,
		"feature_781":  false,
		"feature_782":  false,
		"feature_783":  false,
		"feature_784":  false,
		"feature_785":  false,
		"feature_786":  false,
		"feature_787":  false,
		"feature_788":  false,
		"feature_789":  false,
		"feature_790":  false,
		"feature_791":  false,
		"feature_792":  false,
		"feature_793":  false,
		"feature_794":  false,
		"feature_795":  false,
		"feature_796":  false,
		"feature_797":  false,
		"feature_798":  false,
		"feature_799":  false,
		"feature_800":  false,
		"feature_801":  false,
		"feature_802":  false,
		"feature_803":  false,
		"feature_804":  false,
		"feature_805":  false,
		"feature_806":  false,
		"feature_807":  false,
		"feature_808":  false,
		"feature_809":  false,
		"feature_810":  false,
		"feature_811":  false,
		"feature_812":  false,
		"feature_813":  false,
		"feature_814":  false,
		"feature_815":  false,
		"feature_816":  false,
		"feature_817":  false,
		"feature_818":  false,
		"feature_819":  false,
		"feature_820":  false,
		"feature_821":  false,
		"feature_822":  false,
		"feature_823":  false,
		"feature_824":  false,
		"feature_825":  false,
		"feature_826":  false,
		"feature_827":  false,
		"feature_828":  false,
		"feature_829":  false,
		"feature_830":  false,
		"feature_831":  false,
		"feature_832":  false,
		"feature_833":  false,
		"feature_834":  false,
		"feature_835":  false,
		"feature_836":  false,
		"feature_837":  false,
		"feature_838":  false,
		"feature_839":  false,
		"feature_840":  false,
		"feature_841":  false,
		"feature_842":  false,
		"feature_843":  false,
		"feature_844":  false,
		"feature_845":  false,
		"feature_846":  false,
		"feature_847":  false,
		"feature_848":  false,
		"feature_849":  false,
		"feature_850":  false,
		"feature_851":  false,
		"feature_852":  false,
		"feature_853":  false,
		"feature_854":  false,
		"feature_855":  false,
		"feature_856":  false,
		"feature_857":  false,
		"feature_858":  false,
		"feature_859":  false,
		"feature_860":  false,
		"feature_861":  false,
		"feature_862":  false,
		"feature_863":  false,
		"feature_864":  false,
		"feature_865":  false,
		"feature_866":  false,
		"feature_867":  false,
		"feature_868":  false,
		"feature_869":  false,
		"feature_870":  false,
		"feature_871":  false,
		"feature_872":  false,
		"feature_873":  false,
		"feature_874":  false,
		"feature_875":  false,
		"feature_876":  false,
		"feature_877":  false,
		"feature_878":  false,
		"feature_879":  false,
		"feature_880":  false,
		"feature_881":  false,
		"feature_882":  false,
		"feature_883":  false,
		"feature_884":  false,
		"feature_885":  false,
		"feature_886":  false,
		"feature_887":  false,
		"feature_888":  false,
		"feature_889":  false,
		"feature_890":  false,
		"feature_891":  false,
		"feature_892":  false,
		"feature_893":  false,
		"feature_894":  false,
		"feature_895":  false,
		"feature_896":  false,
		"feature_897":  false,
		"feature_898":  false,
		"feature_899":  false,
		"feature_900":  false,
		"feature_901":  false,
		"feature_902":  false,
		"feature_903":  false,
		"feature_904":  false,
		"feature_905":  false,
		"feature_906":  false,
		"feature_907":  false,
		"feature_908":  false,
		"feature_909":  false,
		"feature_910":  false,
		"feature_911":  false,
		"feature_912":  false,
		"feature_913":  false,
		"feature_914":  false,
		"feature_915":  false,
		"feature_916":  false,
		"feature_917":  false,
		"feature_918":  false,
		"feature_919":  false,
		"feature_920":  false,
		"feature_921":  false,
		"feature_922":  false,
		"feature_923":  false,
		"feature_924":  false,
		"feature_925":  false,
		"feature_926":  false,
		"feature_927":  false,
		"feature_928":  false,
		"feature_929":  false,
		"feature_930":  false,
		"feature_931":  false,
		"feature_932":  false,
		"feature_933":  false,
		"feature_934":  false,
		"feature_935":  false,
		"feature_936":  false,
		"feature_937":  false,
		"feature_938":  false,
		"feature_939":  false,
		"feature_940":  false,
		"feature_941":  false,
		"feature_942":  false,
		"feature_943":  false,
		"feature_944":  false,
		"feature_945":  false,
		"feature_946":  false,
		"feature_947":  false,
		"feature_948":  false,
		"feature_949":  false,
		"feature_950":  false,
		"feature_951":  false,
		"feature_952":  false,
		"feature_953":  false,
		"feature_954":  false,
		"feature_955":  false,
		"feature_956":  false,
		"feature_957":  false,
		"feature_958":  false,
		"feature_959":  false,
		"feature_960":  false,
		"feature_961":  false,
		"feature_962":  false,
		"feature_963":  false,
		"feature_964":  false,
		"feature_965":  false,
		"feature_966":  false,
		"feature_967":  false,
		"feature_968":  false,
		"feature_969":  false,
		"feature_970":  false,
		"feature_971":  false,
		"feature_972":  false,
		"feature_973":  false,
		"feature_974":  false,
		"feature_975":  false,
		"feature_976":  false,
		"feature_977":  false,
		"feature_978":  false,
		"feature_979":  false,
		"feature_980":  false,
		"feature_981":  false,
		"feature_982":  false,
		"feature_983":  false,
		"feature_984":  false,
		"feature_985":  false,
		"feature_986":  false,
		"feature_987":  false,
		"feature_988":  false,
		"feature_989":  false,
		"feature_990":  false,
		"feature_991":  false,
		"feature_992":  false,
		"feature_993":  false,
		"feature_994":  false,
		"feature_995":  false,
		"feature_996":  false,
		"feature_997":  false,
		"feature_998":  false,
		"feature_999":  false,
		"feature_1000": false,
		"feature_1001": false,
		"feature_1002": false,
		"feature_1003": false,
		"feature_1004": false,
		"feature_1005": false,
		"feature_1006": false,
		"feature_1007": false,
		"feature_1008": false,
		"feature_1009": false,
		"feature_1010": false,
		"feature_1011": false,
		"feature_1012": false,
		"feature_1013": false,
		"feature_1014": false,
		"feature_1015": false,
		"feature_1016": false,
		"feature_1017": false,
		"feature_1018": false,
		"feature_1019": false,
		"feature_1020": false,
		"feature_1021": false,
		"feature_1022": false,
		"feature_1023": false,
		"feature_1024": false,
		"feature_1025": false,
		"feature_1026": false,
		"feature_1027": false,
		"feature_1028": false,
		"feature_1029": false,
		"feature_1030": false,
		"feature_1031": false,
		"feature_1032": false,
		"feature_1033": false,
		"feature_1034": false,
		"feature_1035": false,
		"feature_1036": false,
		"feature_1037": false,
		"feature_1038": false,
		"feature_1039": false,
		"feature_1040": false,
		"feature_1041": false,
		"feature_1042": false,
		"feature_1043": false,
		"feature_1044": false,
		"feature_1045": false,
		"feature_1046": false,
		"feature_1047": false,
		"feature_1048": false,
		"feature_1049": false,
		"feature_1050": false,
		"feature_1051": false,
		"feature_1052": false,
		"feature_1053": false,
		"feature_1054": false,
		"feature_1055": false,
		"feature_1056": false,
		"feature_1057": false,
		"feature_1058": false,
		"feature_1059": false,
		"feature_1060": false,
		"feature_1061": false,
		"feature_1062": false,
		"feature_1063": false,
		"feature_1064": false,
		"feature_1065": false,
		"feature_1066": false,
		"feature_1067": false,
		"feature_1068": false,
		"feature_1069": false,
		"feature_1070": false,
		"feature_1071": false,
		"feature_1072": false,
		"feature_1073": false,
		"feature_1074": false,
		"feature_1075": false,
		"feature_1076": false,
		"feature_1077": false,
		"feature_1078": false,
		"feature_1079": false,
		"feature_1080": false,
		"feature_1081": false,
		"feature_1082": false,
		"feature_1083": false,
		"feature_1084": false,
		"feature_1085": false,
		"feature_1086": false,
		"feature_1087": false,
		"feature_1088": false,
		"feature_1089": false,
		"feature_1090": false,
		"feature_1091": false,
		"feature_1092": false,
		"feature_1093": false,
		"feature_1094": false,
		"feature_1095": false,
		"feature_1096": false,
		"feature_1097": false,
		"feature_1098": false,
		"feature_1099": false,
		"feature_1100": false,
		"feature_1101": false,
		"feature_1102": false,
		"feature_1103": false,
		"feature_1104": false,
		"feature_1105": false,
		"feature_1106": false,
		"feature_1107": false,
		"feature_1108": false,
		"feature_1109": false,
		"feature_1110": false,
		"feature_1111": false,
		"feature_1112": false,
		"feature_1113": false,
		"feature_1114": false,
		"feature_1115": false,
		"feature_1116": false,
		"feature_1117": false,
		"feature_1118": false,
		"feature_1119": false,
		"feature_1120": false,
		"feature_1121": false,
		"feature_1122": false,
		"feature_1123": false,
		"feature_1124": false,
		"feature_1125": false,
		"feature_1126": false,
		"feature_1127": false,
		"feature_1128": false,
		"feature_1129": false,
		"feature_1130": false,
		"feature_1131": false,
		"feature_1132": false,
		"feature_1133": false,
		"feature_1134": false,
		"feature_1135": false,
		"feature_1136": false,
		"feature_1137": false,
		"feature_1138": false,
		"feature_1139": false,
		"feature_1140": false,
		"feature_1141": false,
		"feature_1142": false,
		"feature_1143": false,
		"feature_1144": false,
		"feature_1145": false,
		"feature_1146": false,
		"feature_1147": false,
		"feature_1148": false,
		"feature_1149": false,
		"feature_1150": false,
		"feature_1151": false,
		"feature_1152": false,
		"feature_1153": false,
		"feature_1154": false,
		"feature_1155": false,
		"feature_1156": false,
		"feature_1157": false,
		"feature_1158": false,
		"feature_1159": false,
		"feature_1160": false,
		"feature_1161": false,
		"feature_1162": false,
		"feature_1163": false,
		"feature_1164": false,
		"feature_1165": false,
		"feature_1166": false,
		"feature_1167": false,
		"feature_1168": false,
		"feature_1169": false,
		"feature_1170": false,
		"feature_1171": false,
		"feature_1172": false,
		"feature_1173": false,
		"feature_1174": false,
		"feature_1175": false,
		"feature_1176": false,
		"feature_1177": false,
		"feature_1178": false,
		"feature_1179": false,
		"feature_1180": false,
		"feature_1181": false,
		"feature_1182": false,
		"feature_1183": false,
		"feature_1184": false,
		"feature_1185": false,
		"feature_1186": false,
		"feature_1187": false,
		"feature_1188": false,
		"feature_1189": false,
		"feature_1190": false,
		"feature_1191": false,
		"feature_1192": false,
		"feature_1193": false,
		"feature_1194": false,
		"feature_1195": false,
		"feature_1196": false,
		"feature_1197": false,
		"feature_1198": false,
		"feature_1199": false,
		"feature_1200": false,
	}
)

// Simple dumb RW Mutex
type SimpleMutexConfigs struct {
	m       sync.RWMutex
	configs map[string]bool
}

func (c *SimpleMutexConfigs) Update(newConfigs map[string]bool) {
	c.m.Lock()
	defer c.m.Unlock()

	c.configs = map[string]bool{}
	for k, v := range newConfigs {
		c.configs[k] = v
	}
}

func (c *SimpleMutexConfigs) Get(config string) (bool, error) {
	c.m.RLock()
	defer c.m.RUnlock()
	if result, ok := c.configs[config]; ok {
		return result, nil
	}
	return false, errors.New("invalid config")
}

// Pseudo RCU RW Mutex
type PseudoRCUMutexConfigs struct {
	m       sync.RWMutex
	configs map[string]bool
}

func (c *PseudoRCUMutexConfigs) Update(newConfigs map[string]bool) {
	copyConfigs := map[string]bool{}
	for k, v := range newConfigs {
		copyConfigs[k] = v
	}

	c.m.Lock()
	defer c.m.Unlock()
	c.configs = copyConfigs
}

func (c *PseudoRCUMutexConfigs) Get(config string) (bool, error) {
	c.m.RLock()
	defer c.m.RUnlock()
	if result, ok := c.configs[config]; ok {
		return result, nil
	}
	return false, errors.New("invalid config")
}

// Atomic RCU
type AtomicRCUConfigs struct {
	configs atomic.Value
}

func (c *AtomicRCUConfigs) Update(newConfigs map[string]bool) {
	copyConfigs := map[string]bool{}
	for k, v := range newConfigs {
		copyConfigs[k] = v
	}
	c.configs.Store(copyConfigs)
}

func (c *AtomicRCUConfigs) Get(config string) (bool, error) {
	currentConfigs, ok := c.configs.Load().(map[string]bool)
	if !ok {
		return false, errors.New("invalid config state")
	}

	if result, ok := currentConfigs[config]; ok {
		return result, nil
	}
	return false, errors.New("invalid config")
}

type configInterface interface {
	Update(map[string]bool)
	Get(string) (bool, error)
}

func updater(cfg configInterface, c map[string]bool, interval time.Duration, stopChan chan struct{}) {
	ticker := time.NewTicker(interval)
	for {
		select {
		case <-stopChan:
			return
		case <-ticker.C:
			cfg.Update(c)
		}
	}
}

func percentile(values []int64, pct float64) float64 {
	roughIndex := float64(len(values)/100) * pct
	a, b := roughIndex, roughIndex
	if math.Floor(roughIndex) != roughIndex {
		a = math.Floor(roughIndex)
		b = a + 1
	}
	if a >= float64(len(values)) {
		a = float64(len(values)) - 1
		b = a
	}
	return (float64(values[int(a)]) + float64(values[int(b)])) / 2
}

func printPercentiles() {
	for testCase, latencies := range percentiles {
		sort.Slice(latencies, func(i, j int) bool {
			return latencies[i] < latencies[j]
		})
		fmt.Printf("%-*s - \tP50:%-*dP90:%-*dP99:%-*dMax:%-*d\n",
			50, testCase,
			15, int64(percentile(latencies, 50.0)),
			15, int64(percentile(latencies, 90.0)),
			15, int64(percentile(latencies, 99.0)),
			15, int64(percentile(latencies, 100.0)),
		)
	}
}

type latencyMetric struct {
	testCase string
	value    int64
}

func BenchmarkSimpleMutex(b *testing.B) {
	latencyCh := make(chan latencyMetric, 100000000)
	percentiles = map[string][]int64{}
	go func() {
		for {
			metric, ok := <-latencyCh
			if !ok {
				return
			}
			percentiles[metric.testCase] = append(percentiles[metric.testCase], metric.value)
		}
	}()
	c := &SimpleMutexConfigs{}
	wg := &sync.WaitGroup{}
	stopChan := make(chan struct{})
	go updater(c, configs, 5*time.Nanosecond, stopChan)
	b.Run("Update Small Config frequent", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				now := time.Now().UnixNano()
				_, _ = c.Get("feature_0")
				latencyCh <- latencyMetric{"Simple Mutex - Update Small Config frequent", time.Now().UnixNano() - now}
			}()
		}
		wg.Wait()
	})
	close(stopChan)

	c = &SimpleMutexConfigs{}
	stopChan = make(chan struct{})
	go updater(c, configs, 1*time.Second, stopChan)
	b.Run("Update Small Config seldom", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				now := time.Now().UnixNano()
				_, _ = c.Get("feature_0")
				latencyCh <- latencyMetric{"Simple Mutex - Update Small Config seldom", time.Now().UnixNano() - now}
			}()
		}
		wg.Wait()
	})
	close(stopChan)

	c = &SimpleMutexConfigs{}
	stopChan = make(chan struct{})
	go updater(c, largeConfigs, 5*time.Nanosecond, stopChan)
	b.Run("Update Large Config frequent", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				now := time.Now().UnixNano()
				_, _ = c.Get("feature_0")
				latencyCh <- latencyMetric{"Simple Mutex - Update Large Config frequent", time.Now().UnixNano() - now}
			}()
		}
		wg.Wait()
	})
	close(stopChan)

	c = &SimpleMutexConfigs{}
	stopChan = make(chan struct{})
	go updater(c, largeConfigs, 1*time.Second, stopChan)
	b.Run("Update Large Config seldom", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				now := time.Now().UnixNano()
				_, _ = c.Get("feature_0")
				latencyCh <- latencyMetric{"Simple Mutex - Update Large Config seldom", time.Now().UnixNano() - now}
			}()
		}
		wg.Wait()
	})
	close(stopChan)
	close(latencyCh)
	printPercentiles()
}

func BenchmarkPseudoRCU(b *testing.B) {
	latencyCh := make(chan latencyMetric, 100000000)
	percentiles = map[string][]int64{}
	go func() {
		for {
			metric, ok := <-latencyCh
			if !ok {
				return
			}
			percentiles[metric.testCase] = append(percentiles[metric.testCase], metric.value)
		}
	}()
	c := &PseudoRCUMutexConfigs{}
	wg := &sync.WaitGroup{}
	stopChan := make(chan struct{})
	go updater(c, configs, 5*time.Nanosecond, stopChan)
	b.Run("Update Small Config frequent", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				now := time.Now().UnixNano()
				_, _ = c.Get("feature_0")
				latencyCh <- latencyMetric{"Pseudo RCU - Update Small Config frequent", time.Now().UnixNano() - now}
			}()
		}
		wg.Wait()
	})
	close(stopChan)

	c = &PseudoRCUMutexConfigs{}
	stopChan = make(chan struct{})
	go updater(c, configs, 1*time.Second, stopChan)
	b.Run("Update Small Config seldom", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				now := time.Now().UnixNano()
				_, _ = c.Get("feature_0")
				latencyCh <- latencyMetric{"Pseudo RCU - Update Small Config seldom", time.Now().UnixNano() - now}
			}()
		}
		wg.Wait()
	})
	close(stopChan)

	c = &PseudoRCUMutexConfigs{}
	stopChan = make(chan struct{})
	go updater(c, largeConfigs, 5*time.Nanosecond, stopChan)
	b.Run("Update Large Config frequent", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				now := time.Now().UnixNano()
				_, _ = c.Get("feature_0")
				latencyCh <- latencyMetric{"Pseudo RCU - Update Large Config frequent", time.Now().UnixNano() - now}
			}()
		}
		wg.Wait()
	})
	close(stopChan)

	c = &PseudoRCUMutexConfigs{}
	stopChan = make(chan struct{})
	go updater(c, largeConfigs, 1*time.Second, stopChan)
	b.Run("Update Large Config seldom", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				now := time.Now().UnixNano()
				_, _ = c.Get("feature_0")
				latencyCh <- latencyMetric{"Pseudo RCU - Update Large Config seldom", time.Now().UnixNano() - now}
			}()
		}
		wg.Wait()
	})
	close(stopChan)
	close(latencyCh)
	printPercentiles()
}

func BenchmarkAtomicRCU(b *testing.B) {
	latencyCh := make(chan latencyMetric, 100000000)
	percentiles = map[string][]int64{}
	go func() {
		for {
			metric, ok := <-latencyCh
			if !ok {
				return
			}
			percentiles[metric.testCase] = append(percentiles[metric.testCase], metric.value)
		}
	}()
	c := &AtomicRCUConfigs{}
	wg := &sync.WaitGroup{}
	stopChan := make(chan struct{})
	go updater(c, configs, 5*time.Nanosecond, stopChan)
	b.Run("Update Small Config frequent", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				now := time.Now().UnixNano()
				_, _ = c.Get("feature_0")
				latencyCh <- latencyMetric{"Atomic RCU - Update Small Config frequent", time.Now().UnixNano() - now}
			}()
		}
		wg.Wait()
	})
	close(stopChan)

	c = &AtomicRCUConfigs{}
	stopChan = make(chan struct{})
	go updater(c, configs, 1*time.Second, stopChan)
	b.Run("Update Small Config seldom", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				now := time.Now().UnixNano()
				_, _ = c.Get("feature_0")
				latencyCh <- latencyMetric{"Atomic RCU - Update Small Config seldom", time.Now().UnixNano() - now}
			}()
		}
		wg.Wait()
	})
	close(stopChan)

	c = &AtomicRCUConfigs{}
	stopChan = make(chan struct{})
	go updater(c, largeConfigs, 5*time.Nanosecond, stopChan)
	b.Run("Update Large Config frequent", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				now := time.Now().UnixNano()
				_, _ = c.Get("feature_0")
				latencyCh <- latencyMetric{"Atomic RCU - Update Large Config frequent", time.Now().UnixNano() - now}
			}()
		}
		wg.Wait()
	})
	close(stopChan)

	c = &AtomicRCUConfigs{}
	stopChan = make(chan struct{})
	go updater(c, largeConfigs, 1*time.Second, stopChan)
	b.Run("Update Large Config seldom", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				now := time.Now().UnixNano()
				_, _ = c.Get("feature_0")
				latencyCh <- latencyMetric{"Atomic RCU - Update Large Config seldom", time.Now().UnixNano() - now}
			}()
		}
		wg.Wait()
	})
	close(stopChan)
	close(latencyCh)
	printPercentiles()
}
