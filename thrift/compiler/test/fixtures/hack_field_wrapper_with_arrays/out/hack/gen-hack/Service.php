<?hh
/**
 * Autogenerated by Thrift
 *
 * DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING
 *  @generated
 */

/**
 * Original thrift service:-
 * Service
 */
interface ServiceAsyncIf extends \IThriftAsyncIf {
  /**
   * Original thrift definition:-
   * i32
   *   func(1: string arg1,
   *        2: include.MyStruct arg2);
   */
  public function func(string $arg1, ?MyStruct $arg2): Awaitable<int>;
}

/**
 * Original thrift service:-
 * Service
 */
interface ServiceIf extends \IThriftSyncIf {
  /**
   * Original thrift definition:-
   * i32
   *   func(1: string arg1,
   *        2: include.MyStruct arg2);
   */
  public function func(string $arg1, ?MyStruct $arg2): int;
}

/**
 * Original thrift service:-
 * Service
 */
interface ServiceAsyncClientIf extends ServiceAsyncIf {
}

/**
 * Original thrift service:-
 * Service
 */
interface ServiceClientIf extends \IThriftSyncIf {
  /**
   * Original thrift definition:-
   * i32
   *   func(1: string arg1,
   *        2: include.MyStruct arg2);
   */
  public function func(string $arg1, ?MyStruct $arg2): Awaitable<int>;
}

/**
 * Original thrift service:-
 * Service
 */
trait ServiceClientBase {
  require extends \ThriftClientBase;

  /**
   * Original thrift definition:-
   * i32
   *   func(1: string arg1,
   *        2: include.MyStruct arg2);
   */
  public async function func(string $arg1, ?MyStruct $arg2): Awaitable<int> {
    $hh_frame_metadata = $this->getHHFrameMetadata();
    if ($hh_frame_metadata !== null) {
      \HH\set_frame_metadata($hh_frame_metadata);
    }
    $rpc_options = $this->getAndResetOptions() ?? \ThriftClientBase::defaultOptions();
    $args = Service_func_args::fromShape(shape(
      'arg1' => $arg1,
      'arg2' => $arg2,
    ));
    await $this->asyncHandler_->genBefore("Service", "func", $args);
    $currentseqid = $this->sendImplHelper($args, "func", false, "Service" );
    return await $this->genAwaitResponse(Service_func_result::class, "func", false, $currentseqid, $rpc_options);
  }

}

class ServiceAsyncClient extends \ThriftClientBase implements ServiceAsyncClientIf {
  use ServiceClientBase;

}

class ServiceClient extends \ThriftClientBase implements ServiceClientIf {
  use ServiceClientBase;

}

abstract class ServiceAsyncProcessorBase extends \ThriftAsyncProcessor {
  use \GetThriftServiceMetadata;
  abstract const type TThriftIf as ServiceAsyncIf;
  const classname<\IThriftServiceStaticMetadata> SERVICE_METADATA_CLASS = ServiceStaticMetadata::class;
  const string THRIFT_SVC_NAME = 'Service';

  protected async function process_func(int $seqid, \TProtocol $input, \TProtocol $output): Awaitable<void> {
    $handler_ctx = $this->eventHandler_->getHandlerContext('func');
    $reply_type = \TMessageType::REPLY;
    $args = $this->readHelper(Service_func_args::class, $input, 'func', $handler_ctx);
    $result = Service_func_result::withDefaultValues();
    try {
      $this->eventHandler_->preExec($handler_ctx, 'Service', 'func', $args);
      $result->success = await $this->handler->func($args->arg1, $args->arg2);
      $this->eventHandler_->postExec($handler_ctx, 'func', $result);
    } catch (\Exception $ex) {
      $reply_type = \TMessageType::EXCEPTION;
      $this->eventHandler_->handlerError($handler_ctx, 'func', $ex);
      $result = new \TApplicationException($ex->getMessage()."\n".$ex->getTraceAsString());
    }
    $this->writeHelper($result, 'func', $seqid, $handler_ctx, $output, $reply_type);
  }
  protected async function process_getThriftServiceMetadata(int $seqid, \TProtocol $input, \TProtocol $output): Awaitable<void> {
    $this->process_getThriftServiceMetadataHelper($seqid, $input, $output, ServiceStaticMetadata::class);
  }
}
class ServiceAsyncProcessor extends ServiceAsyncProcessorBase {
  const type TThriftIf = ServiceAsyncIf;
}

abstract class ServiceSyncProcessorBase extends \ThriftSyncProcessor {
  use \GetThriftServiceMetadata;
  abstract const type TThriftIf as ServiceIf;
  const classname<\IThriftServiceStaticMetadata> SERVICE_METADATA_CLASS = ServiceStaticMetadata::class;
  const string THRIFT_SVC_NAME = 'Service';

  protected function process_func(int $seqid, \TProtocol $input, \TProtocol $output): void {
    $handler_ctx = $this->eventHandler_->getHandlerContext('func');
    $reply_type = \TMessageType::REPLY;
    $args = $this->readHelper(Service_func_args::class, $input, 'func', $handler_ctx);
    $result = Service_func_result::withDefaultValues();
    try {
      $this->eventHandler_->preExec($handler_ctx, 'Service', 'func', $args);
      $result->success = $this->handler->func($args->arg1, $args->arg2);
      $this->eventHandler_->postExec($handler_ctx, 'func', $result);
    } catch (\Exception $ex) {
      $reply_type = \TMessageType::EXCEPTION;
      $this->eventHandler_->handlerError($handler_ctx, 'func', $ex);
      $result = new \TApplicationException($ex->getMessage()."\n".$ex->getTraceAsString());
    }
    $this->writeHelper($result, 'func', $seqid, $handler_ctx, $output, $reply_type);
  }
  protected function process_getThriftServiceMetadata(int $seqid, \TProtocol $input, \TProtocol $output): void {
    $this->process_getThriftServiceMetadataHelper($seqid, $input, $output, ServiceStaticMetadata::class);
  }
}
class ServiceSyncProcessor extends ServiceSyncProcessorBase {
  const type TThriftIf = ServiceIf;
}
// For backwards compatibility
class ServiceProcessor extends ServiceSyncProcessor {}

// HELPER FUNCTIONS AND STRUCTURES

class Service_func_args implements \IThriftSyncStruct, \IThriftStructMetadata, \IThriftShapishSyncStruct {
  use \ThriftSerializationTrait;

  const \ThriftStructTypes::TSpec SPEC = dict[
    1 => shape(
      'var' => 'arg1',
      'type' => \TType::STRING,
    ),
    2 => shape(
      'var' => 'arg2',
      'type' => \TType::STRUCT,
      'class' => MyStruct::class,
    ),
  ];
  const dict<string, int> FIELDMAP = dict[
    'arg1' => 1,
    'arg2' => 2,
  ];

  const type TConstructorShape = shape(
    ?'arg1' => ?string,
    ?'arg2' => ?MyStruct,
  );

  const type TShape = shape(
    'arg1' => string,
    ?'arg2' => ?MyStruct::TShape,
  );
  const int STRUCTURAL_ID = 6560600906529955702;
  public string $arg1;
  public ?MyStruct $arg2;

  public function __construct(?string $arg1 = null, ?MyStruct $arg2 = null)[] {
    $this->arg1 = $arg1 ?? '';
    $this->arg2 = $arg2;
  }

  public static function withDefaultValues()[]: this {
    return new static();
  }

  public static function fromShape(self::TConstructorShape $shape)[]: this {
    return new static(
      Shapes::idx($shape, 'arg1'),
      Shapes::idx($shape, 'arg2'),
    );
  }

  public static function fromMap_DEPRECATED(@KeyedContainer<string, mixed> $map)[]: this {
    return new static(
      HH\FIXME\UNSAFE_CAST<mixed, string>(idx($map, 'arg1'), 'map value is mixed'),
      HH\FIXME\UNSAFE_CAST<mixed, MyStruct>(idx($map, 'arg2'), 'map value is mixed'),
    );
  }

  public function getName()[]: string {
    return 'Service_func_args';
  }

  public static function getStructMetadata()[]: \tmeta_ThriftStruct {
    return tmeta_ThriftStruct::fromShape(
      shape(
        "name" => "module.func_args",
        "fields" => vec[
          tmeta_ThriftField::fromShape(
            shape(
              "id" => 1,
              "type" => tmeta_ThriftType::fromShape(
                shape(
                  "t_primitive" => tmeta_ThriftPrimitiveType::THRIFT_STRING_TYPE,
                )
              ),
              "name" => "arg1",
            )
          ),
          tmeta_ThriftField::fromShape(
            shape(
              "id" => 2,
              "type" => tmeta_ThriftType::fromShape(
                shape(
                  "t_struct" => tmeta_ThriftStructType::fromShape(
                    shape(
                      "name" => "include.MyStruct",
                    )
                  ),
                )
              ),
              "name" => "arg2",
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

  public static function __stringifyMapKeys<T>(dict<arraykey, T> $m)[]: dict<string, T> {
    return Dict\map_keys($m, $key ==> (string)$key);
  }

  public static function __fromShape(self::TShape $shape)[]: this {
    return new static(
      $shape['arg1'],
      Shapes::idx($shape, 'arg2') |> $$ === null ? null : (MyStruct::__fromShape($$)),
    );
  }

  public function __toShape()[]: self::TShape {
    return shape(
      'arg1' => $this->arg1,
      'arg2' => $this->arg2?->__toShape(),
    );
  }
  public function getInstanceKey()[write_props]: string {
    return \TCompactSerializer::serialize($this);
  }

  public function readFromJson(string $jsonText): void {
    $parsed = json_decode($jsonText, true);

    if ($parsed === null || !($parsed is KeyedContainer<_, _>)) {
      throw new \TProtocolException("Cannot parse the given json string.");
    }

    if (idx($parsed, 'arg1') !== null) {
      $this->arg1 = HH\FIXME\UNSAFE_CAST<mixed, string>($parsed['arg1']);
    }
    if (idx($parsed, 'arg2') !== null) {
      $_tmp0 = \json_encode(HH\FIXME\UNSAFE_CAST<mixed, MyStruct>($parsed['arg2']));
      $_tmp1 = MyStruct::withDefaultValues();
      $_tmp1->readFromJson($_tmp0);
      $this->arg2 = $_tmp1;
    }
  }

}

class Service_func_result extends \ThriftSyncStructWithResult implements \IThriftStructMetadata {
  use \ThriftSerializationTrait;

  const type TResult = int;

  const \ThriftStructTypes::TSpec SPEC = dict[
    0 => shape(
      'var' => 'success',
      'type' => \TType::I32,
    ),
  ];
  const dict<string, int> FIELDMAP = dict[
    'success' => 0,
  ];

  const type TConstructorShape = shape(
    ?'success' => ?this::TResult,
  );

  const int STRUCTURAL_ID = 3865318819874171525;
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
      HH\FIXME\UNSAFE_CAST<mixed, int>(idx($map, 'success'), 'map value is mixed'),
    );
  }

  public function getName()[]: string {
    return 'Service_func_result';
  }

  public static function getStructMetadata()[]: \tmeta_ThriftStruct {
    return tmeta_ThriftStruct::fromShape(
      shape(
        "name" => "module.Service_func_result",
        "fields" => vec[
          tmeta_ThriftField::fromShape(
            shape(
              "id" => 0,
              "type" => tmeta_ThriftType::fromShape(
                shape(
                  "t_primitive" => tmeta_ThriftPrimitiveType::THRIFT_I32_TYPE,
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

  public function readFromJson(string $jsonText): void {
    $parsed = json_decode($jsonText, true);

    if ($parsed === null || !($parsed is KeyedContainer<_, _>)) {
      throw new \TProtocolException("Cannot parse the given json string.");
    }

    if (idx($parsed, 'success') !== null) {
      $_tmp0 = (int)HH\FIXME\UNSAFE_CAST<mixed, int>($parsed['success']);
      if ($_tmp0 > 0x7fffffff) {
        throw new \TProtocolException("number exceeds limit in field");
      } else {
        $this->success = (int)$_tmp0;
      }
    }
  }

}

class ServiceStaticMetadata implements \IThriftServiceStaticMetadata {
  public static function getServiceMetadata()[]: \tmeta_ThriftService {
    return tmeta_ThriftService::fromShape(
      shape(
        "name" => "module.Service",
        "functions" => vec[
          tmeta_ThriftFunction::fromShape(
            shape(
              "name" => "func",
              "return_type" => tmeta_ThriftType::fromShape(
                shape(
                  "t_primitive" => tmeta_ThriftPrimitiveType::THRIFT_I32_TYPE,
                )
              ),
              "arguments" => vec[
                tmeta_ThriftField::fromShape(
                  shape(
                    "id" => 1,
                    "type" => tmeta_ThriftType::fromShape(
                      shape(
                        "t_primitive" => tmeta_ThriftPrimitiveType::THRIFT_STRING_TYPE,
                      )
                    ),
                    "name" => "arg1",
                  )
                ),
                tmeta_ThriftField::fromShape(
                  shape(
                    "id" => 2,
                    "type" => tmeta_ThriftType::fromShape(
                      shape(
                        "t_struct" => tmeta_ThriftStructType::fromShape(
                          shape(
                            "name" => "include.MyStruct",
                          )
                        ),
                      )
                    ),
                    "name" => "arg2",
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
              'include.MyStruct' => MyStruct::getStructMetadata(),
              'include.MyNestedStruct' => MyNestedStruct::getStructMetadata(),
              'include.StructWithWrapper' => \thrift_adapted_types\StructWithWrapper::getStructMetadata(),
              'include.MyStruct' => MyStruct::getStructMetadata(),
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

