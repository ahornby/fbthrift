<?hh
/**
 * Autogenerated by Thrift
 *
 * DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING
 *  @generated
 */

/**
 * Original thrift service:-
 * FooBarBazService
 */
interface FooBarBazServiceAsyncIf extends \IThriftAsyncIf {
  /**
   * Original thrift definition:-
   * void
   *   foo();
   */
  public function foo(): Awaitable<void>;

  /**
   * Original thrift definition:-
   * void
   *   bar();
   */
  public function bar(): Awaitable<void>;

  /**
   * Original thrift definition:-
   * void
   *   baz();
   */
  public function baz(): Awaitable<void>;
}

/**
 * Original thrift service:-
 * FooBarBazService
 */
interface FooBarBazServiceIf extends \IThriftSyncIf {
  /**
   * Original thrift definition:-
   * void
   *   foo();
   */
  public function foo(): void;

  /**
   * Original thrift definition:-
   * void
   *   bar();
   */
  public function bar(): void;

  /**
   * Original thrift definition:-
   * void
   *   baz();
   */
  public function baz(): void;
}

/**
 * Original thrift service:-
 * FooBarBazService
 */
interface FooBarBazServiceAsyncClientIf extends FooBarBazServiceAsyncIf {
}

/**
 * Original thrift service:-
 * FooBarBazService
 */
interface FooBarBazServiceClientIf extends \IThriftSyncIf {
  /**
   * Original thrift definition:-
   * void
   *   foo();
   */
  public function foo(): Awaitable<void>;

  /**
   * Original thrift definition:-
   * void
   *   bar();
   */
  public function bar(): Awaitable<void>;

  /**
   * Original thrift definition:-
   * void
   *   baz();
   */
  public function baz(): Awaitable<void>;
}

/**
 * Original thrift service:-
 * FooBarBazService
 */
trait FooBarBazServiceClientBase {
  require extends \ThriftClientBase;

  /**
   * Original thrift definition:-
   * void
   *   foo();
   */
  public async function foo(): Awaitable<void> {
    $hh_frame_metadata = $this->getHHFrameMetadata();
    if ($hh_frame_metadata !== null) {
      \HH\set_frame_metadata($hh_frame_metadata);
    }
    $rpc_options = $this->getAndResetOptions() ?? \ThriftClientBase::defaultOptions();
    $args = FooBarBazService_foo_args::withDefaultValues();
    await $this->asyncHandler_->genBefore(FooBarBazServiceStaticMetadata::THRIFT_SVC_NAME, "foo", $args);
    $currentseqid = $this->sendImplHelper($args, "foo", false, FooBarBazServiceStaticMetadata::THRIFT_SVC_NAME );
    await $this->genAwaitResponse(FooBarBazService_foo_result::class, "foo", true, $currentseqid, $rpc_options);
  }

  /**
   * Original thrift definition:-
   * void
   *   bar();
   */
  public async function bar(): Awaitable<void> {
    $hh_frame_metadata = $this->getHHFrameMetadata();
    if ($hh_frame_metadata !== null) {
      \HH\set_frame_metadata($hh_frame_metadata);
    }
    $rpc_options = $this->getAndResetOptions() ?? \ThriftClientBase::defaultOptions();
    $args = FooBarBazService_bar_args::withDefaultValues();
    await $this->asyncHandler_->genBefore(FooBarBazServiceStaticMetadata::THRIFT_SVC_NAME, "bar", $args);
    $currentseqid = $this->sendImplHelper($args, "bar", false, FooBarBazServiceStaticMetadata::THRIFT_SVC_NAME );
    await $this->genAwaitResponse(FooBarBazService_bar_result::class, "bar", true, $currentseqid, $rpc_options);
  }

  /**
   * Original thrift definition:-
   * void
   *   baz();
   */
  public async function baz(): Awaitable<void> {
    $hh_frame_metadata = $this->getHHFrameMetadata();
    if ($hh_frame_metadata !== null) {
      \HH\set_frame_metadata($hh_frame_metadata);
    }
    $rpc_options = $this->getAndResetOptions() ?? \ThriftClientBase::defaultOptions();
    $args = FooBarBazService_baz_args::withDefaultValues();
    await $this->asyncHandler_->genBefore(FooBarBazServiceStaticMetadata::THRIFT_SVC_NAME, "baz", $args);
    $currentseqid = $this->sendImplHelper($args, "baz", false, FooBarBazServiceStaticMetadata::THRIFT_SVC_NAME );
    await $this->genAwaitResponse(FooBarBazService_baz_result::class, "baz", true, $currentseqid, $rpc_options);
  }

}

class FooBarBazServiceAsyncClient extends \ThriftClientBase implements FooBarBazServiceAsyncClientIf {
  use FooBarBazServiceClientBase;

  const string THRIFT_SVC_NAME = FooBarBazServiceStaticMetadata::THRIFT_SVC_NAME;

}

class FooBarBazServiceClient extends \ThriftClientBase implements FooBarBazServiceClientIf {
  use FooBarBazServiceClientBase;

  const string THRIFT_SVC_NAME = FooBarBazServiceStaticMetadata::THRIFT_SVC_NAME;

}

abstract class FooBarBazServiceAsyncProcessorBase extends \ThriftAsyncProcessor {
  use \GetThriftServiceMetadata;
  abstract const type TThriftIf as FooBarBazServiceAsyncIf;
  const classname<\IThriftServiceStaticMetadata> SERVICE_METADATA_CLASS = FooBarBazServiceStaticMetadata::class;
  const string THRIFT_SVC_NAME = FooBarBazServiceStaticMetadata::THRIFT_SVC_NAME;

  protected async function process_foo(int $seqid, \TProtocol $input, \TProtocol $output): Awaitable<void> {
    $handler_ctx = $this->eventHandler_->getHandlerContext('foo');
    $reply_type = \TMessageType::REPLY;
    $args = $this->readHelper(FooBarBazService_foo_args::class, $input, 'foo', $handler_ctx);
    $result = FooBarBazService_foo_result::withDefaultValues();
    try {
      $this->eventHandler_->preExec($handler_ctx, 'FooBarBazService', 'foo', $args);
      await $this->handler->foo();
      $this->eventHandler_->postExec($handler_ctx, 'foo', $result);
    } catch (\Exception $ex) {
      $reply_type = \TMessageType::EXCEPTION;
      $this->eventHandler_->handlerError($handler_ctx, 'foo', $ex);
      $result = new \TApplicationException($ex->getMessage()."\n".$ex->getTraceAsString());
    }
    $this->writeHelper($result, 'foo', $seqid, $handler_ctx, $output, $reply_type);
  }
  protected async function process_bar(int $seqid, \TProtocol $input, \TProtocol $output): Awaitable<void> {
    $handler_ctx = $this->eventHandler_->getHandlerContext('bar');
    $reply_type = \TMessageType::REPLY;
    $args = $this->readHelper(FooBarBazService_bar_args::class, $input, 'bar', $handler_ctx);
    $result = FooBarBazService_bar_result::withDefaultValues();
    try {
      $this->eventHandler_->preExec($handler_ctx, 'FooBarBazService', 'bar', $args);
      await $this->handler->bar();
      $this->eventHandler_->postExec($handler_ctx, 'bar', $result);
    } catch (\Exception $ex) {
      $reply_type = \TMessageType::EXCEPTION;
      $this->eventHandler_->handlerError($handler_ctx, 'bar', $ex);
      $result = new \TApplicationException($ex->getMessage()."\n".$ex->getTraceAsString());
    }
    $this->writeHelper($result, 'bar', $seqid, $handler_ctx, $output, $reply_type);
  }
  protected async function process_baz(int $seqid, \TProtocol $input, \TProtocol $output): Awaitable<void> {
    $handler_ctx = $this->eventHandler_->getHandlerContext('baz');
    $reply_type = \TMessageType::REPLY;
    $args = $this->readHelper(FooBarBazService_baz_args::class, $input, 'baz', $handler_ctx);
    $result = FooBarBazService_baz_result::withDefaultValues();
    try {
      $this->eventHandler_->preExec($handler_ctx, 'FooBarBazService', 'baz', $args);
      await $this->handler->baz();
      $this->eventHandler_->postExec($handler_ctx, 'baz', $result);
    } catch (\Exception $ex) {
      $reply_type = \TMessageType::EXCEPTION;
      $this->eventHandler_->handlerError($handler_ctx, 'baz', $ex);
      $result = new \TApplicationException($ex->getMessage()."\n".$ex->getTraceAsString());
    }
    $this->writeHelper($result, 'baz', $seqid, $handler_ctx, $output, $reply_type);
  }
  protected async function process_getThriftServiceMetadata(int $seqid, \TProtocol $input, \TProtocol $output): Awaitable<void> {
    $this->process_getThriftServiceMetadataHelper($seqid, $input, $output, FooBarBazServiceStaticMetadata::class);
  }
}
class FooBarBazServiceAsyncProcessor extends FooBarBazServiceAsyncProcessorBase {
  const type TThriftIf = FooBarBazServiceAsyncIf;
}

abstract class FooBarBazServiceSyncProcessorBase extends \ThriftSyncProcessor {
  use \GetThriftServiceMetadata;
  abstract const type TThriftIf as FooBarBazServiceIf;
  const classname<\IThriftServiceStaticMetadata> SERVICE_METADATA_CLASS = FooBarBazServiceStaticMetadata::class;
  const string THRIFT_SVC_NAME = FooBarBazServiceStaticMetadata::THRIFT_SVC_NAME;

  protected function process_foo(int $seqid, \TProtocol $input, \TProtocol $output): void {
    $handler_ctx = $this->eventHandler_->getHandlerContext('foo');
    $reply_type = \TMessageType::REPLY;
    $args = $this->readHelper(FooBarBazService_foo_args::class, $input, 'foo', $handler_ctx);
    $result = FooBarBazService_foo_result::withDefaultValues();
    try {
      $this->eventHandler_->preExec($handler_ctx, 'FooBarBazService', 'foo', $args);
      $this->handler->foo();
      $this->eventHandler_->postExec($handler_ctx, 'foo', $result);
    } catch (\Exception $ex) {
      $reply_type = \TMessageType::EXCEPTION;
      $this->eventHandler_->handlerError($handler_ctx, 'foo', $ex);
      $result = new \TApplicationException($ex->getMessage()."\n".$ex->getTraceAsString());
    }
    $this->writeHelper($result, 'foo', $seqid, $handler_ctx, $output, $reply_type);
  }
  protected function process_bar(int $seqid, \TProtocol $input, \TProtocol $output): void {
    $handler_ctx = $this->eventHandler_->getHandlerContext('bar');
    $reply_type = \TMessageType::REPLY;
    $args = $this->readHelper(FooBarBazService_bar_args::class, $input, 'bar', $handler_ctx);
    $result = FooBarBazService_bar_result::withDefaultValues();
    try {
      $this->eventHandler_->preExec($handler_ctx, 'FooBarBazService', 'bar', $args);
      $this->handler->bar();
      $this->eventHandler_->postExec($handler_ctx, 'bar', $result);
    } catch (\Exception $ex) {
      $reply_type = \TMessageType::EXCEPTION;
      $this->eventHandler_->handlerError($handler_ctx, 'bar', $ex);
      $result = new \TApplicationException($ex->getMessage()."\n".$ex->getTraceAsString());
    }
    $this->writeHelper($result, 'bar', $seqid, $handler_ctx, $output, $reply_type);
  }
  protected function process_baz(int $seqid, \TProtocol $input, \TProtocol $output): void {
    $handler_ctx = $this->eventHandler_->getHandlerContext('baz');
    $reply_type = \TMessageType::REPLY;
    $args = $this->readHelper(FooBarBazService_baz_args::class, $input, 'baz', $handler_ctx);
    $result = FooBarBazService_baz_result::withDefaultValues();
    try {
      $this->eventHandler_->preExec($handler_ctx, 'FooBarBazService', 'baz', $args);
      $this->handler->baz();
      $this->eventHandler_->postExec($handler_ctx, 'baz', $result);
    } catch (\Exception $ex) {
      $reply_type = \TMessageType::EXCEPTION;
      $this->eventHandler_->handlerError($handler_ctx, 'baz', $ex);
      $result = new \TApplicationException($ex->getMessage()."\n".$ex->getTraceAsString());
    }
    $this->writeHelper($result, 'baz', $seqid, $handler_ctx, $output, $reply_type);
  }
  protected function process_getThriftServiceMetadata(int $seqid, \TProtocol $input, \TProtocol $output): void {
    $this->process_getThriftServiceMetadataHelper($seqid, $input, $output, FooBarBazServiceStaticMetadata::class);
  }
}
class FooBarBazServiceSyncProcessor extends FooBarBazServiceSyncProcessorBase {
  const type TThriftIf = FooBarBazServiceIf;
}
// For backwards compatibility
class FooBarBazServiceProcessor extends FooBarBazServiceSyncProcessor {}

// HELPER FUNCTIONS AND STRUCTURES

class FooBarBazService_foo_args implements \IThriftSyncStruct, \IThriftStructMetadata, \IThriftShapishSyncStruct {
  use \ThriftSerializationTrait;

  const \ThriftStructTypes::TSpec SPEC = dict[
  ];
  const dict<string, int> FIELDMAP = dict[
  ];

  const type TConstructorShape = shape(
  );

  const type TShape = shape(
    ...
  );
  const int STRUCTURAL_ID = 957977401221134810;

  public function __construct()[] {
  }

  public static function withDefaultValues()[]: this {
    return new static();
  }

  public static function fromShape(self::TConstructorShape $shape)[]: this {
    return new static(
    );
  }

  public function getName()[]: string {
    return 'FooBarBazService_foo_args';
  }

  public static function getStructMetadata()[]: \tmeta_ThriftStruct {
    return tmeta_ThriftStruct::fromShape(
      shape(
        "name" => "module.foo_args",
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

  public static function __fromShape(self::TShape $shape)[]: this {
    return new static(
    );
  }

  public function __toShape()[]: self::TShape {
    return shape(
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

  }

}

class FooBarBazService_foo_result extends \ThriftSyncStructWithoutResult implements \IThriftStructMetadata {
  use \ThriftSerializationTrait;

  const \ThriftStructTypes::TSpec SPEC = dict[
  ];
  const dict<string, int> FIELDMAP = dict[
  ];

  const type TConstructorShape = shape(
  );

  const int STRUCTURAL_ID = 957977401221134810;

  public function __construct()[] {
  }

  public static function withDefaultValues()[]: this {
    return new static();
  }

  public static function fromShape(self::TConstructorShape $shape)[]: this {
    return new static(
    );
  }

  public function getName()[]: string {
    return 'FooBarBazService_foo_result';
  }

  public static function getStructMetadata()[]: \tmeta_ThriftStruct {
    return tmeta_ThriftStruct::fromShape(
      shape(
        "name" => "module.FooBarBazService_foo_result",
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

  }

}

class FooBarBazService_bar_args implements \IThriftSyncStruct, \IThriftStructMetadata, \IThriftShapishSyncStruct {
  use \ThriftSerializationTrait;

  const \ThriftStructTypes::TSpec SPEC = dict[
  ];
  const dict<string, int> FIELDMAP = dict[
  ];

  const type TConstructorShape = shape(
  );

  const type TShape = shape(
    ...
  );
  const int STRUCTURAL_ID = 957977401221134810;

  public function __construct()[] {
  }

  public static function withDefaultValues()[]: this {
    return new static();
  }

  public static function fromShape(self::TConstructorShape $shape)[]: this {
    return new static(
    );
  }

  public function getName()[]: string {
    return 'FooBarBazService_bar_args';
  }

  public static function getStructMetadata()[]: \tmeta_ThriftStruct {
    return tmeta_ThriftStruct::fromShape(
      shape(
        "name" => "module.bar_args",
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

  public static function __fromShape(self::TShape $shape)[]: this {
    return new static(
    );
  }

  public function __toShape()[]: self::TShape {
    return shape(
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

  }

}

class FooBarBazService_bar_result extends \ThriftSyncStructWithoutResult implements \IThriftStructMetadata {
  use \ThriftSerializationTrait;

  const \ThriftStructTypes::TSpec SPEC = dict[
  ];
  const dict<string, int> FIELDMAP = dict[
  ];

  const type TConstructorShape = shape(
  );

  const int STRUCTURAL_ID = 957977401221134810;

  public function __construct()[] {
  }

  public static function withDefaultValues()[]: this {
    return new static();
  }

  public static function fromShape(self::TConstructorShape $shape)[]: this {
    return new static(
    );
  }

  public function getName()[]: string {
    return 'FooBarBazService_bar_result';
  }

  public static function getStructMetadata()[]: \tmeta_ThriftStruct {
    return tmeta_ThriftStruct::fromShape(
      shape(
        "name" => "module.FooBarBazService_bar_result",
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

  }

}

class FooBarBazService_baz_args implements \IThriftSyncStruct, \IThriftStructMetadata, \IThriftShapishSyncStruct {
  use \ThriftSerializationTrait;

  const \ThriftStructTypes::TSpec SPEC = dict[
  ];
  const dict<string, int> FIELDMAP = dict[
  ];

  const type TConstructorShape = shape(
  );

  const type TShape = shape(
    ...
  );
  const int STRUCTURAL_ID = 957977401221134810;

  public function __construct()[] {
  }

  public static function withDefaultValues()[]: this {
    return new static();
  }

  public static function fromShape(self::TConstructorShape $shape)[]: this {
    return new static(
    );
  }

  public function getName()[]: string {
    return 'FooBarBazService_baz_args';
  }

  public static function getStructMetadata()[]: \tmeta_ThriftStruct {
    return tmeta_ThriftStruct::fromShape(
      shape(
        "name" => "module.baz_args",
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

  public static function __fromShape(self::TShape $shape)[]: this {
    return new static(
    );
  }

  public function __toShape()[]: self::TShape {
    return shape(
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

  }

}

class FooBarBazService_baz_result extends \ThriftSyncStructWithoutResult implements \IThriftStructMetadata {
  use \ThriftSerializationTrait;

  const \ThriftStructTypes::TSpec SPEC = dict[
  ];
  const dict<string, int> FIELDMAP = dict[
  ];

  const type TConstructorShape = shape(
  );

  const int STRUCTURAL_ID = 957977401221134810;

  public function __construct()[] {
  }

  public static function withDefaultValues()[]: this {
    return new static();
  }

  public static function fromShape(self::TConstructorShape $shape)[]: this {
    return new static(
    );
  }

  public function getName()[]: string {
    return 'FooBarBazService_baz_result';
  }

  public static function getStructMetadata()[]: \tmeta_ThriftStruct {
    return tmeta_ThriftStruct::fromShape(
      shape(
        "name" => "module.FooBarBazService_baz_result",
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

  }

}

class FooBarBazServiceStaticMetadata implements \IThriftServiceStaticMetadata {
  const string THRIFT_SVC_NAME = 'FooBarBazService';

  public static function getServiceMetadata()[]: \tmeta_ThriftService {
    return tmeta_ThriftService::fromShape(
      shape(
        "name" => "module.FooBarBazService",
        "functions" => vec[
          tmeta_ThriftFunction::fromShape(
            shape(
              "name" => "foo",
              "return_type" => tmeta_ThriftType::fromShape(
                shape(
                  "t_primitive" => tmeta_ThriftPrimitiveType::THRIFT_VOID_TYPE,
                )
              ),
            )
          ),
          tmeta_ThriftFunction::fromShape(
            shape(
              "name" => "bar",
              "return_type" => tmeta_ThriftType::fromShape(
                shape(
                  "t_primitive" => tmeta_ThriftPrimitiveType::THRIFT_VOID_TYPE,
                )
              ),
            )
          ),
          tmeta_ThriftFunction::fromShape(
            shape(
              "name" => "baz",
              "return_type" => tmeta_ThriftType::fromShape(
                shape(
                  "t_primitive" => tmeta_ThriftPrimitiveType::THRIFT_VOID_TYPE,
                )
              ),
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
        'foo' => dict[
          '\facebook\thrift\annotation\go\Name' => \facebook\thrift\annotation\go\Name::fromShape(
            shape(
              "name" => "FooStructured",
            )
          ),
        ],
        'bar' => dict[
          '\facebook\thrift\annotation\go\Name' => \facebook\thrift\annotation\go\Name::fromShape(
            shape(
              "name" => "BarNonStructured",
            )
          ),
        ],
      ],
    );
  }
}

