<?hh
/**
 * Autogenerated by Thrift
 *
 * DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING
 *  @generated
 */

/**
 * Original thrift service:-
 * Bar
 */
interface BarAsyncIf extends \IThriftAsyncIf {
  /**
   * Original thrift definition:-
   * string
   *   baz(1: set<i32> a,
   *       2: list<map<i32, set<string>>> b,
   *       3: i64 c,
   *       4: Foo d,
   *       5: i64 e);
   */
  public function baz(?dict<int, bool> $a, ?KeyedContainer<int, KeyedContainer<int, dict<string, bool>>> $b, ?int $c, ?Foo $d, ?int $e): Awaitable<string>;
}

/**
 * Original thrift service:-
 * Bar
 */
interface BarIf extends \IThriftSyncIf {
  /**
   * Original thrift definition:-
   * string
   *   baz(1: set<i32> a,
   *       2: list<map<i32, set<string>>> b,
   *       3: i64 c,
   *       4: Foo d,
   *       5: i64 e);
   */
  public function baz(?dict<int, bool> $a, ?KeyedContainer<int, KeyedContainer<int, dict<string, bool>>> $b, ?int $c, ?Foo $d, ?int $e): string;
}

/**
 * Original thrift service:-
 * Bar
 */
interface BarAsyncClientIf extends BarAsyncIf {
}

/**
 * Original thrift service:-
 * Bar
 */
interface BarClientIf extends \IThriftSyncIf {
  /**
   * Original thrift definition:-
   * string
   *   baz(1: set<i32> a,
   *       2: list<map<i32, set<string>>> b,
   *       3: i64 c,
   *       4: Foo d,
   *       5: i64 e);
   */
  public function baz(?dict<int, bool> $a, ?KeyedContainer<int, KeyedContainer<int, dict<string, bool>>> $b, ?int $c, ?Foo $d, ?int $e): Awaitable<string>;
}

/**
 * Original thrift service:-
 * Bar
 */
trait BarClientBase {
  require extends \ThriftClientBase;

  /**
   * Original thrift definition:-
   * string
   *   baz(1: set<i32> a,
   *       2: list<map<i32, set<string>>> b,
   *       3: i64 c,
   *       4: Foo d,
   *       5: i64 e);
   */
  public async function baz(?dict<int, bool> $a, ?KeyedContainer<int, KeyedContainer<int, dict<string, bool>>> $b, ?int $c, ?Foo $d, ?int $e): Awaitable<string> {
    $hh_frame_metadata = $this->getHHFrameMetadata();
    if ($hh_frame_metadata !== null) {
      \HH\set_frame_metadata($hh_frame_metadata);
    }
    $rpc_options = $this->getAndResetOptions() ?? \ThriftClientBase::defaultOptions();
    $args = Bar_baz_args::fromShape(shape(
      'a' => $a === null ? null : $a,
      'b' => $b === null ? null : Vec\map($b, 
        $_val0 ==> dict($_val0)
      ),
      'c' => $c === null ? null : $c,
      'd' => $d === null ? null : $d,
      'e' => $e === null ? null : $e,
    ));
    await $this->asyncHandler_->genBefore("Bar", "baz", $args);
    $currentseqid = $this->sendImplHelper($args, "baz", false, "Bar" );
    return await $this->genAwaitResponse(Bar_baz_result::class, "baz", false, $currentseqid, $rpc_options);
  }

  /**
   * Original thrift definition:-
   * string
   *   baz(1: set<i32> a,
   *       2: list<map<i32, set<string>>> b,
   *       3: i64 c,
   *       4: Foo d,
   *       5: i64 e);
   */
  public async function baz__LEGACY_ARRAYS(?dict<int, bool> $a, ?KeyedContainer<int, KeyedContainer<int, dict<string, bool>>> $b, ?int $c, ?Foo $d, ?int $e): Awaitable<string> {
    $hh_frame_metadata = $this->getHHFrameMetadata();
    if ($hh_frame_metadata !== null) {
      \HH\set_frame_metadata($hh_frame_metadata);
    }
    $rpc_options = $this->getAndResetOptions() ?? \ThriftClientBase::defaultOptions();
    $args = Bar_baz_args::fromShape(shape(
      'a' => $a === null ? null : $a,
      'b' => $b === null ? null : Vec\map($b, 
        $_val0 ==> dict($_val0)
      ),
      'c' => $c === null ? null : $c,
      'd' => $d === null ? null : $d,
      'e' => $e === null ? null : $e,
    ));
    await $this->asyncHandler_->genBefore("Bar", "baz", $args);
    $currentseqid = $this->sendImplHelper($args, "baz", false, "Bar" );
    return await $this->genAwaitResponse(Bar_baz_result::class, "baz", false, $currentseqid, $rpc_options, shape('read_options' => \THRIFT_MARK_LEGACY_ARRAYS));
  }

}

class BarAsyncClient extends \ThriftClientBase implements BarAsyncClientIf {
  use BarClientBase;

}

class BarClient extends \ThriftClientBase implements BarClientIf {
  use BarClientBase;

}

abstract class BarAsyncProcessorBase extends \ThriftAsyncProcessor {
  use \GetThriftServiceMetadata;
  abstract const type TThriftIf as BarAsyncIf;
  const classname<\IThriftServiceStaticMetadata> SERVICE_METADATA_CLASS = BarStaticMetadata::class;
  const string THRIFT_SVC_NAME = 'Bar';

  protected async function process_baz(int $seqid, \TProtocol $input, \TProtocol $output): Awaitable<void> {
    $handler_ctx = $this->eventHandler_->getHandlerContext('baz');
    $reply_type = \TMessageType::REPLY;
    $args = $this->readHelper(Bar_baz_args::class, $input, 'baz', $handler_ctx);
    $result = Bar_baz_result::withDefaultValues();
    try {
      $this->eventHandler_->preExec($handler_ctx, 'Bar', 'baz', $args);
      $result->success = await $this->handler->baz($args->a, $args->b, $args->c, $args->d, $args->e);
      $this->eventHandler_->postExec($handler_ctx, 'baz', $result);
    } catch (\Exception $ex) {
      $reply_type = \TMessageType::EXCEPTION;
      $this->eventHandler_->handlerError($handler_ctx, 'baz', $ex);
      $result = new \TApplicationException($ex->getMessage()."\n".$ex->getTraceAsString());
    }
    $this->writeHelper($result, 'baz', $seqid, $handler_ctx, $output, $reply_type);
  }
  protected async function process_getThriftServiceMetadata(int $seqid, \TProtocol $input, \TProtocol $output): Awaitable<void> {
    $this->process_getThriftServiceMetadataHelper($seqid, $input, $output, BarStaticMetadata::class);
  }
}
class BarAsyncProcessor extends BarAsyncProcessorBase {
  const type TThriftIf = BarAsyncIf;
}

abstract class BarSyncProcessorBase extends \ThriftSyncProcessor {
  use \GetThriftServiceMetadata;
  abstract const type TThriftIf as BarIf;
  const classname<\IThriftServiceStaticMetadata> SERVICE_METADATA_CLASS = BarStaticMetadata::class;
  const string THRIFT_SVC_NAME = 'Bar';

  protected function process_baz(int $seqid, \TProtocol $input, \TProtocol $output): void {
    $handler_ctx = $this->eventHandler_->getHandlerContext('baz');
    $reply_type = \TMessageType::REPLY;
    $args = $this->readHelper(Bar_baz_args::class, $input, 'baz', $handler_ctx);
    $result = Bar_baz_result::withDefaultValues();
    try {
      $this->eventHandler_->preExec($handler_ctx, 'Bar', 'baz', $args);
      $result->success = $this->handler->baz($args->a, $args->b, $args->c, $args->d, $args->e);
      $this->eventHandler_->postExec($handler_ctx, 'baz', $result);
    } catch (\Exception $ex) {
      $reply_type = \TMessageType::EXCEPTION;
      $this->eventHandler_->handlerError($handler_ctx, 'baz', $ex);
      $result = new \TApplicationException($ex->getMessage()."\n".$ex->getTraceAsString());
    }
    $this->writeHelper($result, 'baz', $seqid, $handler_ctx, $output, $reply_type);
  }
  protected function process_getThriftServiceMetadata(int $seqid, \TProtocol $input, \TProtocol $output): void {
    $this->process_getThriftServiceMetadataHelper($seqid, $input, $output, BarStaticMetadata::class);
  }
}
class BarSyncProcessor extends BarSyncProcessorBase {
  const type TThriftIf = BarIf;
}
// For backwards compatibility
class BarProcessor extends BarSyncProcessor {}

// HELPER FUNCTIONS AND STRUCTURES

class Bar_baz_args implements \IThriftSyncStruct, \IThriftStructMetadata {
  use \ThriftSerializationTrait;

  const \ThriftStructTypes::TSpec SPEC = dict[
    1 => shape(
      'var' => 'a',
      'type' => \TType::SET,
      'etype' => \TType::I32,
      'elem' => shape(
        'type' => \TType::I32,
      ),
      'format' => 'array',
    ),
    2 => shape(
      'var' => 'b',
      'type' => \TType::LST,
      'etype' => \TType::MAP,
      'elem' => shape(
        'type' => \TType::MAP,
        'ktype' => \TType::I32,
        'vtype' => \TType::SET,
        'key' => shape(
          'type' => \TType::I32,
        ),
        'val' => shape(
          'type' => \TType::SET,
          'etype' => \TType::STRING,
          'elem' => shape(
            'type' => \TType::STRING,
          ),
          'format' => 'array',
        ),
        'format' => 'array',
      ),
      'format' => 'array',
    ),
    3 => shape(
      'var' => 'c',
      'type' => \TType::I64,
    ),
    4 => shape(
      'var' => 'd',
      'type' => \TType::STRUCT,
      'class' => Foo::class,
    ),
    5 => shape(
      'var' => 'e',
      'type' => \TType::I64,
    ),
  ];
  const dict<string, int> FIELDMAP = dict[
    'a' => 1,
    'b' => 2,
    'c' => 3,
    'd' => 4,
    'e' => 5,
  ];

  const type TConstructorShape = shape(
    ?'a' => ?dict<int, bool>,
    ?'b' => ?vec<dict<int, dict<string, bool>>>,
    ?'c' => ?int,
    ?'d' => ?Foo,
    ?'e' => ?int,
  );

  const int STRUCTURAL_ID = 7865027497865509792;
  public ?dict<int, bool> $a;
  public ?vec<dict<int, dict<string, bool>>> $b;
  public ?int $c;
  public ?Foo $d;
  public ?int $e;

  public function __construct(?dict<int, bool> $a = null, ?vec<dict<int, dict<string, bool>>> $b = null, ?int $c = null, ?Foo $d = null, ?int $e = null)[] {
    $this->a = $a;
    $this->b = $b;
    $this->c = $c;
    $this->d = $d;
    $this->e = $e ?? 4;
  }

  public static function withDefaultValues()[]: this {
    return new static();
  }

  public static function fromShape(self::TConstructorShape $shape)[]: this {
    return new static(
      Shapes::idx($shape, 'a'),
      Shapes::idx($shape, 'b'),
      Shapes::idx($shape, 'c'),
      Shapes::idx($shape, 'd'),
      Shapes::idx($shape, 'e'),
    );
  }

  public static function fromMap_DEPRECATED(@KeyedContainer<string, mixed> $map)[]: this {
    return new static(
      HH\FIXME\UNSAFE_CAST<mixed, dict<int, bool>>(idx($map, 'a'), 'map value is mixed'),
      HH\FIXME\UNSAFE_CAST<mixed, vec<dict<int, dict<string, bool>>>>(idx($map, 'b'), 'map value is mixed'),
      HH\FIXME\UNSAFE_CAST<mixed, int>(idx($map, 'c'), 'map value is mixed'),
      HH\FIXME\UNSAFE_CAST<mixed, Foo>(idx($map, 'd'), 'map value is mixed'),
      HH\FIXME\UNSAFE_CAST<mixed, int>(idx($map, 'e'), 'map value is mixed'),
    );
  }

  public function getName()[]: string {
    return 'Bar_baz_args';
  }

  public static function getStructMetadata()[]: \tmeta_ThriftStruct {
    return tmeta_ThriftStruct::fromShape(
      shape(
        "name" => "module.baz_args",
        "fields" => vec[
          tmeta_ThriftField::fromShape(
            shape(
              "id" => 1,
              "type" => tmeta_ThriftType::fromShape(
                shape(
                  "t_set" => tmeta_ThriftSetType::fromShape(
                    shape(
                      "valueType" => tmeta_ThriftType::fromShape(
                        shape(
                          "t_primitive" => tmeta_ThriftPrimitiveType::THRIFT_I32_TYPE,
                        )
                      ),
                    )
                  ),
                )
              ),
              "name" => "a",
            )
          ),
          tmeta_ThriftField::fromShape(
            shape(
              "id" => 2,
              "type" => tmeta_ThriftType::fromShape(
                shape(
                  "t_list" => tmeta_ThriftListType::fromShape(
                    shape(
                      "valueType" => tmeta_ThriftType::fromShape(
                        shape(
                          "t_map" => tmeta_ThriftMapType::fromShape(
                            shape(
                              "keyType" => tmeta_ThriftType::fromShape(
                                shape(
                                  "t_primitive" => tmeta_ThriftPrimitiveType::THRIFT_I32_TYPE,
                                )
                              ),
                              "valueType" => tmeta_ThriftType::fromShape(
                                shape(
                                  "t_set" => tmeta_ThriftSetType::fromShape(
                                    shape(
                                      "valueType" => tmeta_ThriftType::fromShape(
                                        shape(
                                          "t_primitive" => tmeta_ThriftPrimitiveType::THRIFT_STRING_TYPE,
                                        )
                                      ),
                                    )
                                  ),
                                )
                              ),
                            )
                          ),
                        )
                      ),
                    )
                  ),
                )
              ),
              "name" => "b",
            )
          ),
          tmeta_ThriftField::fromShape(
            shape(
              "id" => 3,
              "type" => tmeta_ThriftType::fromShape(
                shape(
                  "t_primitive" => tmeta_ThriftPrimitiveType::THRIFT_I64_TYPE,
                )
              ),
              "name" => "c",
            )
          ),
          tmeta_ThriftField::fromShape(
            shape(
              "id" => 4,
              "type" => tmeta_ThriftType::fromShape(
                shape(
                  "t_struct" => tmeta_ThriftStructType::fromShape(
                    shape(
                      "name" => "module.Foo",
                    )
                  ),
                )
              ),
              "name" => "d",
            )
          ),
          tmeta_ThriftField::fromShape(
            shape(
              "id" => 5,
              "type" => tmeta_ThriftType::fromShape(
                shape(
                  "t_primitive" => tmeta_ThriftPrimitiveType::THRIFT_I64_TYPE,
                )
              ),
              "name" => "e",
            )
          ),
        ],
        "is_union" => false,
      )
    );
  }

  public static function getAllStructuredAnnotations()[write_props]: \TStructAnnotations {
    return shape(
      'struct' => dict[],
      'fields' => dict[
      ],
    );
  }

  public function getInstanceKey()[write_props]: string {
    return \TCompactSerializer::serialize($this);
  }

}

class Bar_baz_result extends \ThriftSyncStructWithResult implements \IThriftStructMetadata {
  use \ThriftSerializationTrait;

  const type TResult = string;

  const \ThriftStructTypes::TSpec SPEC = dict[
    0 => shape(
      'var' => 'success',
      'type' => \TType::STRING,
    ),
  ];
  const dict<string, int> FIELDMAP = dict[
    'success' => 0,
  ];

  const type TConstructorShape = shape(
    ?'success' => ?this::TResult,
  );

  const int STRUCTURAL_ID = 1365128170602685579;
  public ?this::TResult $success;

  public function __construct(?this::TResult $success = null)[] {
    $this->success = $success;
  }

  public static function withDefaultValues()[]: this {
    return new static();
  }

  public static function fromShape(self::TConstructorShape $shape)[]: this {
    return new static(
      Shapes::idx($shape, 'success'),
    );
  }

  public static function fromMap_DEPRECATED(@KeyedContainer<string, mixed> $map)[]: this {
    return new static(
      HH\FIXME\UNSAFE_CAST<mixed, string>(idx($map, 'success'), 'map value is mixed'),
    );
  }

  public function getName()[]: string {
    return 'Bar_baz_result';
  }

  public static function getStructMetadata()[]: \tmeta_ThriftStruct {
    return tmeta_ThriftStruct::fromShape(
      shape(
        "name" => "module.Bar_baz_result",
        "fields" => vec[
          tmeta_ThriftField::fromShape(
            shape(
              "id" => 0,
              "type" => tmeta_ThriftType::fromShape(
                shape(
                  "t_primitive" => tmeta_ThriftPrimitiveType::THRIFT_STRING_TYPE,
                )
              ),
              "name" => "success",
            )
          ),
        ],
        "is_union" => false,
      )
    );
  }

  public static function getAllStructuredAnnotations()[write_props]: \TStructAnnotations {
    return shape(
      'struct' => dict[],
      'fields' => dict[
      ],
    );
  }

  public function getInstanceKey()[write_props]: string {
    return \TCompactSerializer::serialize($this);
  }

}

class BarStaticMetadata implements \IThriftServiceStaticMetadata {
  public static function getServiceMetadata()[]: \tmeta_ThriftService {
    return tmeta_ThriftService::fromShape(
      shape(
        "name" => "module.Bar",
        "functions" => vec[
          tmeta_ThriftFunction::fromShape(
            shape(
              "name" => "baz",
              "return_type" => tmeta_ThriftType::fromShape(
                shape(
                  "t_primitive" => tmeta_ThriftPrimitiveType::THRIFT_STRING_TYPE,
                )
              ),
              "arguments" => vec[
                tmeta_ThriftField::fromShape(
                  shape(
                    "id" => 1,
                    "type" => tmeta_ThriftType::fromShape(
                      shape(
                        "t_set" => tmeta_ThriftSetType::fromShape(
                          shape(
                            "valueType" => tmeta_ThriftType::fromShape(
                              shape(
                                "t_primitive" => tmeta_ThriftPrimitiveType::THRIFT_I32_TYPE,
                              )
                            ),
                          )
                        ),
                      )
                    ),
                    "name" => "a",
                  )
                ),
                tmeta_ThriftField::fromShape(
                  shape(
                    "id" => 2,
                    "type" => tmeta_ThriftType::fromShape(
                      shape(
                        "t_list" => tmeta_ThriftListType::fromShape(
                          shape(
                            "valueType" => tmeta_ThriftType::fromShape(
                              shape(
                                "t_map" => tmeta_ThriftMapType::fromShape(
                                  shape(
                                    "keyType" => tmeta_ThriftType::fromShape(
                                      shape(
                                        "t_primitive" => tmeta_ThriftPrimitiveType::THRIFT_I32_TYPE,
                                      )
                                    ),
                                    "valueType" => tmeta_ThriftType::fromShape(
                                      shape(
                                        "t_set" => tmeta_ThriftSetType::fromShape(
                                          shape(
                                            "valueType" => tmeta_ThriftType::fromShape(
                                              shape(
                                                "t_primitive" => tmeta_ThriftPrimitiveType::THRIFT_STRING_TYPE,
                                              )
                                            ),
                                          )
                                        ),
                                      )
                                    ),
                                  )
                                ),
                              )
                            ),
                          )
                        ),
                      )
                    ),
                    "name" => "b",
                  )
                ),
                tmeta_ThriftField::fromShape(
                  shape(
                    "id" => 3,
                    "type" => tmeta_ThriftType::fromShape(
                      shape(
                        "t_primitive" => tmeta_ThriftPrimitiveType::THRIFT_I64_TYPE,
                      )
                    ),
                    "name" => "c",
                  )
                ),
                tmeta_ThriftField::fromShape(
                  shape(
                    "id" => 4,
                    "type" => tmeta_ThriftType::fromShape(
                      shape(
                        "t_struct" => tmeta_ThriftStructType::fromShape(
                          shape(
                            "name" => "module.Foo",
                          )
                        ),
                      )
                    ),
                    "name" => "d",
                  )
                ),
                tmeta_ThriftField::fromShape(
                  shape(
                    "id" => 5,
                    "type" => tmeta_ThriftType::fromShape(
                      shape(
                        "t_primitive" => tmeta_ThriftPrimitiveType::THRIFT_I64_TYPE,
                      )
                    ),
                    "name" => "e",
                  )
                ),
              ],
            )
          ),
        ],
      )
    );
  }

  public static function getServiceMetadataResponse()[]: \tmeta_ThriftServiceMetadataResponse {
    return \tmeta_ThriftServiceMetadataResponse::fromShape(
      shape(
        'context' => \tmeta_ThriftServiceContext::fromShape(
          shape(
            'service_info' => self::getServiceMetadata(),
            'module' => \tmeta_ThriftModuleContext::fromShape(
              shape(
                'name' => 'module',
              )
            ),
          )
        ),
        'metadata' => \tmeta_ThriftMetadata::fromShape(
          shape(
            'enums' => dict[
            ],
            'structs' => dict[
              'module.Foo' => Foo::getStructMetadata(),
            ],
            'exceptions' => dict[
            ],
            'services' => dict[
            ],
          )
        ),
      )
    );
  }

  public static function getAllStructuredAnnotations()[write_props]: \TServiceAnnotations {
    return shape(
      'service' => dict[],
      'functions' => dict[
      ],
    );
  }
}

