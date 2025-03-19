<?hh
/**
 * Autogenerated by Thrift
 *
 * DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING
 *  @generated
 */

namespace hack\fixtures;

module hack.module.internal;
/**
 * Original thrift service:-
 * TestServiceWithMethodAnnotation
 */
interface TestServiceWithMethodAnnotationAsyncIf extends \IThriftAsyncIf {
  /**
   * Original thrift definition:-
   * i32
   *   testMethodWithAnnotation();
   */
  internal function testMethodWithAnnotation(): Awaitable<int>;

  /**
   * Original thrift definition:-
   * void
   *   testMethodWithoutAnnotation();
   */
  public function testMethodWithoutAnnotation(): Awaitable<void>;
}

/**
 * Original thrift service:-
 * TestServiceWithMethodAnnotation
 */
interface TestServiceWithMethodAnnotationIf extends \IThriftSyncIf {
  /**
   * Original thrift definition:-
   * i32
   *   testMethodWithAnnotation();
   */
  internal function testMethodWithAnnotation(): int;

  /**
   * Original thrift definition:-
   * void
   *   testMethodWithoutAnnotation();
   */
  public function testMethodWithoutAnnotation(): void;
}

/**
 * Original thrift service:-
 * TestServiceWithMethodAnnotation
 */
interface TestServiceWithMethodAnnotationAsyncClientIf extends TestServiceWithMethodAnnotationAsyncIf {
}

/**
 * Original thrift service:-
 * TestServiceWithMethodAnnotation
 */
interface TestServiceWithMethodAnnotationClientIf extends \IThriftSyncIf {
  /**
   * Original thrift definition:-
   * i32
   *   testMethodWithAnnotation();
   */
  internal function testMethodWithAnnotation(): Awaitable<int>;

  /**
   * Original thrift definition:-
   * void
   *   testMethodWithoutAnnotation();
   */
  public function testMethodWithoutAnnotation(): Awaitable<void>;
}

/**
 * Original thrift service:-
 * TestServiceWithMethodAnnotation
 */
internal trait TestServiceWithMethodAnnotationClientBase {
  require extends \ThriftClientBase;

  /**
   * Original thrift definition:-
   * i32
   *   testMethodWithAnnotation();
   */
  internal async function testMethodWithAnnotation(): Awaitable<int> {
    $hh_frame_metadata = $this->getHHFrameMetadata();
    if ($hh_frame_metadata !== null) {
      \HH\set_frame_metadata($hh_frame_metadata);
    }
    $rpc_options = $this->getAndResetOptions() ?? \ThriftClientBase::defaultOptions();
    $args = \hack\fixtures\TestServiceWithMethodAnnotation_testMethodWithAnnotation_args::withDefaultValues();
    await $this->asyncHandler_->genBefore(TestServiceWithMethodAnnotationStaticMetadata::THRIFT_SVC_NAME, "testMethodWithAnnotation", $args);
    $currentseqid = $this->sendImplHelper($args, "testMethodWithAnnotation", false, TestServiceWithMethodAnnotationStaticMetadata::THRIFT_SVC_NAME );
    return (await $this->genAwaitResponse(\hack\fixtures\TestServiceWithMethodAnnotation_testMethodWithAnnotation_result::class, "testMethodWithAnnotation", false, $currentseqid, $rpc_options))[0];
  }

  /**
   * Original thrift definition:-
   * void
   *   testMethodWithoutAnnotation();
   */
  public async function testMethodWithoutAnnotation(): Awaitable<void> {
    $hh_frame_metadata = $this->getHHFrameMetadata();
    if ($hh_frame_metadata !== null) {
      \HH\set_frame_metadata($hh_frame_metadata);
    }
    $rpc_options = $this->getAndResetOptions() ?? \ThriftClientBase::defaultOptions();
    $args = \hack\fixtures\TestServiceWithMethodAnnotation_testMethodWithoutAnnotation_args::withDefaultValues();
    await $this->asyncHandler_->genBefore(TestServiceWithMethodAnnotationStaticMetadata::THRIFT_SVC_NAME, "testMethodWithoutAnnotation", $args);
    $currentseqid = $this->sendImplHelper($args, "testMethodWithoutAnnotation", false, TestServiceWithMethodAnnotationStaticMetadata::THRIFT_SVC_NAME );
    await $this->genAwaitResponse(\hack\fixtures\TestServiceWithMethodAnnotation_testMethodWithoutAnnotation_result::class, "testMethodWithoutAnnotation", true, $currentseqid, $rpc_options);
  }

}

class TestServiceWithMethodAnnotationAsyncClient extends \ThriftClientBase implements TestServiceWithMethodAnnotationAsyncClientIf {
  use TestServiceWithMethodAnnotationClientBase;

  const string THRIFT_SVC_NAME = TestServiceWithMethodAnnotationStaticMetadata::THRIFT_SVC_NAME;

}

class TestServiceWithMethodAnnotationClient extends \ThriftClientBase implements TestServiceWithMethodAnnotationClientIf {
  use TestServiceWithMethodAnnotationClientBase;

  const string THRIFT_SVC_NAME = TestServiceWithMethodAnnotationStaticMetadata::THRIFT_SVC_NAME;

}

// HELPER FUNCTIONS AND STRUCTURES

class TestServiceWithMethodAnnotation_testMethodWithAnnotation_args implements \IThriftSyncStruct, \IThriftStructMetadata {
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
    return 'TestServiceWithMethodAnnotation_testMethodWithAnnotation_args';
  }

  public static function getStructMetadata()[]: \tmeta_ThriftStruct {
    return \tmeta_ThriftStruct::fromShape(
      shape(
        "name" => "module.testMethodWithAnnotation_args",
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

class TestServiceWithMethodAnnotation_testMethodWithAnnotation_result extends \ThriftSyncStructWithResult implements \IThriftStructMetadata {
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

  public function getName()[]: string {
    return 'TestServiceWithMethodAnnotation_testMethodWithAnnotation_result';
  }

  public static function getStructMetadata()[]: \tmeta_ThriftStruct {
    return \tmeta_ThriftStruct::fromShape(
      shape(
        "name" => "module.TestServiceWithMethodAnnotation_testMethodWithAnnotation_result",
        "fields" => vec[
          \tmeta_ThriftField::fromShape(
            shape(
              "id" => 0,
              "type" => \tmeta_ThriftType::fromShape(
                shape(
                  "t_primitive" => \tmeta_ThriftPrimitiveType::THRIFT_I32_TYPE,
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

class TestServiceWithMethodAnnotation_testMethodWithoutAnnotation_args implements \IThriftSyncStruct, \IThriftStructMetadata {
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
    return 'TestServiceWithMethodAnnotation_testMethodWithoutAnnotation_args';
  }

  public static function getStructMetadata()[]: \tmeta_ThriftStruct {
    return \tmeta_ThriftStruct::fromShape(
      shape(
        "name" => "module.testMethodWithoutAnnotation_args",
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

class TestServiceWithMethodAnnotation_testMethodWithoutAnnotation_result extends \ThriftSyncStructWithoutResult implements \IThriftStructMetadata {
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
    return 'TestServiceWithMethodAnnotation_testMethodWithoutAnnotation_result';
  }

  public static function getStructMetadata()[]: \tmeta_ThriftStruct {
    return \tmeta_ThriftStruct::fromShape(
      shape(
        "name" => "module.TestServiceWithMethodAnnotation_testMethodWithoutAnnotation_result",
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

class TestServiceWithMethodAnnotationStaticMetadata implements \IThriftServiceStaticMetadata {
  const string THRIFT_SVC_NAME = 'TestServiceWithMethodAnnotation';

  public static function getServiceMetadata()[]: \tmeta_ThriftService {
    return \tmeta_ThriftService::fromShape(
      shape(
        "name" => "module.TestServiceWithMethodAnnotation",
        "functions" => vec[
          \tmeta_ThriftFunction::fromShape(
            shape(
              "name" => "testMethodWithAnnotation",
              "return_type" => \tmeta_ThriftType::fromShape(
                shape(
                  "t_primitive" => \tmeta_ThriftPrimitiveType::THRIFT_I32_TYPE,
                )
              ),
            )
          ),
          \tmeta_ThriftFunction::fromShape(
            shape(
              "name" => "testMethodWithoutAnnotation",
              "return_type" => \tmeta_ThriftType::fromShape(
                shape(
                  "t_primitive" => \tmeta_ThriftPrimitiveType::THRIFT_VOID_TYPE,
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
        'testMethodWithAnnotation' => dict[
          '\facebook\thrift\annotation\hack\ModuleInternal' => \facebook\thrift\annotation\hack\ModuleInternal::fromShape(
            shape(
            )
          ),
        ],
      ],
    );
  }
}

