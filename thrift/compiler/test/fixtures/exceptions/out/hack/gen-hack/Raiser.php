<?hh
/**
 * Autogenerated by Thrift
 *
 * DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING
 *  @generated
 */

/**
 * Original thrift service:-
 * Raiser
 */
interface RaiserAsyncIf extends \IThriftAsyncIf {
  /**
   * Original thrift definition:-
   * void
   *   doBland();
   */
  public function doBland(): Awaitable<void>;

  /**
   * Original thrift definition:-
   * void
   *   doRaise()
   *   throws (1: Banal b,
   *           2: Fiery f,
   *           3: Serious s);
   */
  public function doRaise(): Awaitable<void>;

  /**
   * Original thrift definition:-
   * string
   *   get200();
   */
  public function get200(): Awaitable<string>;

  /**
   * Original thrift definition:-
   * string
   *   get500()
   *   throws (1: Fiery f,
   *           2: Banal b,
   *           3: Serious s);
   */
  public function get500(): Awaitable<string>;
}

/**
 * Original thrift service:-
 * Raiser
 */
interface RaiserIf extends \IThriftSyncIf {
  /**
   * Original thrift definition:-
   * void
   *   doBland();
   */
  public function doBland(): void;

  /**
   * Original thrift definition:-
   * void
   *   doRaise()
   *   throws (1: Banal b,
   *           2: Fiery f,
   *           3: Serious s);
   */
  public function doRaise(): void;

  /**
   * Original thrift definition:-
   * string
   *   get200();
   */
  public function get200(): string;

  /**
   * Original thrift definition:-
   * string
   *   get500()
   *   throws (1: Fiery f,
   *           2: Banal b,
   *           3: Serious s);
   */
  public function get500(): string;
}

/**
 * Original thrift service:-
 * Raiser
 */
interface RaiserAsyncClientIf extends RaiserAsyncIf {
}

/**
 * Original thrift service:-
 * Raiser
 */
interface RaiserClientIf extends \IThriftSyncIf {
  /**
   * Original thrift definition:-
   * void
   *   doBland();
   */
  public function doBland(): Awaitable<void>;

  /**
   * Original thrift definition:-
   * void
   *   doRaise()
   *   throws (1: Banal b,
   *           2: Fiery f,
   *           3: Serious s);
   */
  public function doRaise(): Awaitable<void>;

  /**
   * Original thrift definition:-
   * string
   *   get200();
   */
  public function get200(): Awaitable<string>;

  /**
   * Original thrift definition:-
   * string
   *   get500()
   *   throws (1: Fiery f,
   *           2: Banal b,
   *           3: Serious s);
   */
  public function get500(): Awaitable<string>;
}

/**
 * Original thrift service:-
 * Raiser
 */
trait RaiserClientBase {
  require extends \ThriftClientBase;

  /**
   * Original thrift definition:-
   * void
   *   doBland();
   */
  public async function doBland(): Awaitable<void> {
    $hh_frame_metadata = $this->getHHFrameMetadata();
    if ($hh_frame_metadata !== null) {
      \HH\set_frame_metadata($hh_frame_metadata);
    }
    $rpc_options = $this->getAndResetOptions() ?? \ThriftClientBase::defaultOptions();
    $args = Raiser_doBland_args::withDefaultValues();
    await $this->asyncHandler_->genBefore(RaiserStaticMetadata::THRIFT_SVC_NAME, "doBland", $args);
    $currentseqid = $this->sendImplHelper($args, "doBland", false, RaiserStaticMetadata::THRIFT_SVC_NAME );
    await $this->genAwaitResponse(Raiser_doBland_result::class, "doBland", true, $currentseqid, $rpc_options);
  }

  /**
   * Original thrift definition:-
   * void
   *   doRaise()
   *   throws (1: Banal b,
   *           2: Fiery f,
   *           3: Serious s);
   */
  public async function doRaise(): Awaitable<void> {
    $hh_frame_metadata = $this->getHHFrameMetadata();
    if ($hh_frame_metadata !== null) {
      \HH\set_frame_metadata($hh_frame_metadata);
    }
    $rpc_options = $this->getAndResetOptions() ?? \ThriftClientBase::defaultOptions();
    $args = Raiser_doRaise_args::withDefaultValues();
    await $this->asyncHandler_->genBefore(RaiserStaticMetadata::THRIFT_SVC_NAME, "doRaise", $args);
    $currentseqid = $this->sendImplHelper($args, "doRaise", false, RaiserStaticMetadata::THRIFT_SVC_NAME );
    await $this->genAwaitResponse(Raiser_doRaise_result::class, "doRaise", true, $currentseqid, $rpc_options);
  }

  /**
   * Original thrift definition:-
   * string
   *   get200();
   */
  public async function get200(): Awaitable<string> {
    $hh_frame_metadata = $this->getHHFrameMetadata();
    if ($hh_frame_metadata !== null) {
      \HH\set_frame_metadata($hh_frame_metadata);
    }
    $rpc_options = $this->getAndResetOptions() ?? \ThriftClientBase::defaultOptions();
    $args = Raiser_get200_args::withDefaultValues();
    await $this->asyncHandler_->genBefore(RaiserStaticMetadata::THRIFT_SVC_NAME, "get200", $args);
    $currentseqid = $this->sendImplHelper($args, "get200", false, RaiserStaticMetadata::THRIFT_SVC_NAME );
    return (await $this->genAwaitResponse(Raiser_get200_result::class, "get200", false, $currentseqid, $rpc_options))[0];
  }

  /**
   * Original thrift definition:-
   * string
   *   get500()
   *   throws (1: Fiery f,
   *           2: Banal b,
   *           3: Serious s);
   */
  public async function get500(): Awaitable<string> {
    $hh_frame_metadata = $this->getHHFrameMetadata();
    if ($hh_frame_metadata !== null) {
      \HH\set_frame_metadata($hh_frame_metadata);
    }
    $rpc_options = $this->getAndResetOptions() ?? \ThriftClientBase::defaultOptions();
    $args = Raiser_get500_args::withDefaultValues();
    await $this->asyncHandler_->genBefore(RaiserStaticMetadata::THRIFT_SVC_NAME, "get500", $args);
    $currentseqid = $this->sendImplHelper($args, "get500", false, RaiserStaticMetadata::THRIFT_SVC_NAME );
    return (await $this->genAwaitResponse(Raiser_get500_result::class, "get500", false, $currentseqid, $rpc_options))[0];
  }

}

class RaiserAsyncClient extends \ThriftClientBase implements RaiserAsyncClientIf {
  use RaiserClientBase;

  const string THRIFT_SVC_NAME = RaiserStaticMetadata::THRIFT_SVC_NAME;

}

class RaiserClient extends \ThriftClientBase implements RaiserClientIf {
  use RaiserClientBase;

  const string THRIFT_SVC_NAME = RaiserStaticMetadata::THRIFT_SVC_NAME;

}

// HELPER FUNCTIONS AND STRUCTURES

class Raiser_doBland_args implements \IThriftSyncStruct, \IThriftStructMetadata {
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
    return 'Raiser_doBland_args';
  }

  public static function getStructMetadata()[]: \tmeta_ThriftStruct {
    return tmeta_ThriftStruct::fromShape(
      shape(
        "name" => "module.doBland_args",
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

class Raiser_doBland_result extends \ThriftSyncStructWithoutResult implements \IThriftStructMetadata {
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
    return 'Raiser_doBland_result';
  }

  public static function getStructMetadata()[]: \tmeta_ThriftStruct {
    return tmeta_ThriftStruct::fromShape(
      shape(
        "name" => "module.Raiser_doBland_result",
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

class Raiser_doRaise_args implements \IThriftSyncStruct, \IThriftStructMetadata {
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
    return 'Raiser_doRaise_args';
  }

  public static function getStructMetadata()[]: \tmeta_ThriftStruct {
    return tmeta_ThriftStruct::fromShape(
      shape(
        "name" => "module.doRaise_args",
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

class Raiser_doRaise_result extends \ThriftSyncStructWithoutResult implements \IThriftStructMetadata {
  use \ThriftSerializationTrait;

  const \ThriftStructTypes::TSpec SPEC = dict[
    1 => shape(
      'var' => 'b',
      'type' => \TType::STRUCT,
      'class' => Banal::class,
    ),
    2 => shape(
      'var' => 'f',
      'type' => \TType::STRUCT,
      'class' => Fiery::class,
    ),
    3 => shape(
      'var' => 's',
      'type' => \TType::STRUCT,
      'class' => Serious::class,
    ),
  ];
  const dict<string, int> FIELDMAP = dict[
    'b' => 1,
    'f' => 2,
    's' => 3,
  ];

  const type TConstructorShape = shape(
    ?'b' => ?Banal,
    ?'f' => ?Fiery,
    ?'s' => ?Serious,
  );

  const int STRUCTURAL_ID = 1991614683033939916;
  public ?Banal $b;
  public ?Fiery $f;
  public ?Serious $s;

  public function __construct(?Banal $b = null, ?Fiery $f = null, ?Serious $s = null)[] {
    $this->b = $b;
    $this->f = $f;
    $this->s = $s;
  }

  public static function withDefaultValues()[]: this {
    return new static();
  }

  public static function fromShape(self::TConstructorShape $shape)[]: this {
    return new static(
      Shapes::idx($shape, 'b'),
      Shapes::idx($shape, 'f'),
      Shapes::idx($shape, 's'),
    );
  }

  public function getName()[]: string {
    return 'Raiser_doRaise_result';
  }

  public static function getStructMetadata()[]: \tmeta_ThriftStruct {
    return tmeta_ThriftStruct::fromShape(
      shape(
        "name" => "module.Raiser_doRaise_result",
        "fields" => vec[
          tmeta_ThriftField::fromShape(
            shape(
              "id" => 1,
              "type" => tmeta_ThriftType::fromShape(
                shape(
                  "t_struct" => tmeta_ThriftStructType::fromShape(
                    shape(
                      "name" => "module.Banal",
                    )
                  ),
                )
              ),
              "name" => "b",
            )
          ),
          tmeta_ThriftField::fromShape(
            shape(
              "id" => 2,
              "type" => tmeta_ThriftType::fromShape(
                shape(
                  "t_struct" => tmeta_ThriftStructType::fromShape(
                    shape(
                      "name" => "module.Fiery",
                    )
                  ),
                )
              ),
              "name" => "f",
            )
          ),
          tmeta_ThriftField::fromShape(
            shape(
              "id" => 3,
              "type" => tmeta_ThriftType::fromShape(
                shape(
                  "t_struct" => tmeta_ThriftStructType::fromShape(
                    shape(
                      "name" => "module.Serious",
                    )
                  ),
                )
              ),
              "name" => "s",
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

  public function checkForException(): ?\TException {
    if ($this->b !== null) {
      return $this->b;
    }
    if ($this->f !== null) {
      return $this->f;
    }
    if ($this->s !== null) {
      return $this->s;
    }
    return null;
  }
  
  public function setException(\Exception $e): bool {
    if ($e is Banal) {
      $this->b = $e;
      return true;
    }
    if ($e is Fiery) {
      $this->f = $e;
      return true;
    }
    if ($e is Serious) {
      $this->s = $e;
      return true;
    }
    return false;
  }
}

class Raiser_get200_args implements \IThriftSyncStruct, \IThriftStructMetadata {
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
    return 'Raiser_get200_args';
  }

  public static function getStructMetadata()[]: \tmeta_ThriftStruct {
    return tmeta_ThriftStruct::fromShape(
      shape(
        "name" => "module.get200_args",
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

class Raiser_get200_result extends \ThriftSyncStructWithResult implements \IThriftStructMetadata {
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

  public function getName()[]: string {
    return 'Raiser_get200_result';
  }

  public static function getStructMetadata()[]: \tmeta_ThriftStruct {
    return tmeta_ThriftStruct::fromShape(
      shape(
        "name" => "module.Raiser_get200_result",
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

class Raiser_get500_args implements \IThriftSyncStruct, \IThriftStructMetadata {
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
    return 'Raiser_get500_args';
  }

  public static function getStructMetadata()[]: \tmeta_ThriftStruct {
    return tmeta_ThriftStruct::fromShape(
      shape(
        "name" => "module.get500_args",
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

class Raiser_get500_result extends \ThriftSyncStructWithResult implements \IThriftStructMetadata {
  use \ThriftSerializationTrait;

  const type TResult = string;

  const \ThriftStructTypes::TSpec SPEC = dict[
    0 => shape(
      'var' => 'success',
      'type' => \TType::STRING,
    ),
    1 => shape(
      'var' => 'f',
      'type' => \TType::STRUCT,
      'class' => Fiery::class,
    ),
    2 => shape(
      'var' => 'b',
      'type' => \TType::STRUCT,
      'class' => Banal::class,
    ),
    3 => shape(
      'var' => 's',
      'type' => \TType::STRUCT,
      'class' => Serious::class,
    ),
  ];
  const dict<string, int> FIELDMAP = dict[
    'success' => 0,
    'f' => 1,
    'b' => 2,
    's' => 3,
  ];

  const type TConstructorShape = shape(
    ?'success' => ?this::TResult,
    ?'f' => ?Fiery,
    ?'b' => ?Banal,
    ?'s' => ?Serious,
  );

  const int STRUCTURAL_ID = 6147773747560615508;
  public ?this::TResult $success;
  public ?Fiery $f;
  public ?Banal $b;
  public ?Serious $s;

  public function __construct(?this::TResult $success = null, ?Fiery $f = null, ?Banal $b = null, ?Serious $s = null)[] {
    $this->success = $success;
    $this->f = $f;
    $this->b = $b;
    $this->s = $s;
  }

  public static function withDefaultValues()[]: this {
    return new static();
  }

  public static function fromShape(self::TConstructorShape $shape)[]: this {
    return new static(
      Shapes::idx($shape, 'success'),
      Shapes::idx($shape, 'f'),
      Shapes::idx($shape, 'b'),
      Shapes::idx($shape, 's'),
    );
  }

  public function getName()[]: string {
    return 'Raiser_get500_result';
  }

  public static function getStructMetadata()[]: \tmeta_ThriftStruct {
    return tmeta_ThriftStruct::fromShape(
      shape(
        "name" => "module.Raiser_get500_result",
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
          tmeta_ThriftField::fromShape(
            shape(
              "id" => 1,
              "type" => tmeta_ThriftType::fromShape(
                shape(
                  "t_struct" => tmeta_ThriftStructType::fromShape(
                    shape(
                      "name" => "module.Fiery",
                    )
                  ),
                )
              ),
              "name" => "f",
            )
          ),
          tmeta_ThriftField::fromShape(
            shape(
              "id" => 2,
              "type" => tmeta_ThriftType::fromShape(
                shape(
                  "t_struct" => tmeta_ThriftStructType::fromShape(
                    shape(
                      "name" => "module.Banal",
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
                  "t_struct" => tmeta_ThriftStructType::fromShape(
                    shape(
                      "name" => "module.Serious",
                    )
                  ),
                )
              ),
              "name" => "s",
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

  public function checkForException(): ?\TException {
    if ($this->f !== null) {
      return $this->f;
    }
    if ($this->b !== null) {
      return $this->b;
    }
    if ($this->s !== null) {
      return $this->s;
    }
    return null;
  }
  
  public function setException(\Exception $e): bool {
    if ($e is Fiery) {
      $this->f = $e;
      return true;
    }
    if ($e is Banal) {
      $this->b = $e;
      return true;
    }
    if ($e is Serious) {
      $this->s = $e;
      return true;
    }
    return false;
  }
}

class RaiserStaticMetadata implements \IThriftServiceStaticMetadata {
  const string THRIFT_SVC_NAME = 'Raiser';

  public static function getServiceMetadata()[]: \tmeta_ThriftService {
    return tmeta_ThriftService::fromShape(
      shape(
        "name" => "module.Raiser",
        "functions" => vec[
          tmeta_ThriftFunction::fromShape(
            shape(
              "name" => "doBland",
              "return_type" => tmeta_ThriftType::fromShape(
                shape(
                  "t_primitive" => tmeta_ThriftPrimitiveType::THRIFT_VOID_TYPE,
                )
              ),
            )
          ),
          tmeta_ThriftFunction::fromShape(
            shape(
              "name" => "doRaise",
              "return_type" => tmeta_ThriftType::fromShape(
                shape(
                  "t_primitive" => tmeta_ThriftPrimitiveType::THRIFT_VOID_TYPE,
                )
              ),
              "exceptions" => vec[
                tmeta_ThriftField::fromShape(
                  shape(
                    "id" => 1,
                    "type" => tmeta_ThriftType::fromShape(
                      shape(
                        "t_struct" => tmeta_ThriftStructType::fromShape(
                          shape(
                            "name" => "module.Banal",
                          )
                        ),
                      )
                    ),
                    "name" => "b",
                  )
                ),
                tmeta_ThriftField::fromShape(
                  shape(
                    "id" => 2,
                    "type" => tmeta_ThriftType::fromShape(
                      shape(
                        "t_struct" => tmeta_ThriftStructType::fromShape(
                          shape(
                            "name" => "module.Fiery",
                          )
                        ),
                      )
                    ),
                    "name" => "f",
                  )
                ),
                tmeta_ThriftField::fromShape(
                  shape(
                    "id" => 3,
                    "type" => tmeta_ThriftType::fromShape(
                      shape(
                        "t_struct" => tmeta_ThriftStructType::fromShape(
                          shape(
                            "name" => "module.Serious",
                          )
                        ),
                      )
                    ),
                    "name" => "s",
                  )
                ),
              ],
            )
          ),
          tmeta_ThriftFunction::fromShape(
            shape(
              "name" => "get200",
              "return_type" => tmeta_ThriftType::fromShape(
                shape(
                  "t_primitive" => tmeta_ThriftPrimitiveType::THRIFT_STRING_TYPE,
                )
              ),
            )
          ),
          tmeta_ThriftFunction::fromShape(
            shape(
              "name" => "get500",
              "return_type" => tmeta_ThriftType::fromShape(
                shape(
                  "t_primitive" => tmeta_ThriftPrimitiveType::THRIFT_STRING_TYPE,
                )
              ),
              "exceptions" => vec[
                tmeta_ThriftField::fromShape(
                  shape(
                    "id" => 1,
                    "type" => tmeta_ThriftType::fromShape(
                      shape(
                        "t_struct" => tmeta_ThriftStructType::fromShape(
                          shape(
                            "name" => "module.Fiery",
                          )
                        ),
                      )
                    ),
                    "name" => "f",
                  )
                ),
                tmeta_ThriftField::fromShape(
                  shape(
                    "id" => 2,
                    "type" => tmeta_ThriftType::fromShape(
                      shape(
                        "t_struct" => tmeta_ThriftStructType::fromShape(
                          shape(
                            "name" => "module.Banal",
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
                        "t_struct" => tmeta_ThriftStructType::fromShape(
                          shape(
                            "name" => "module.Serious",
                          )
                        ),
                      )
                    ),
                    "name" => "s",
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
            ],
            'exceptions' => dict[
              'module.Banal' => Banal::getExceptionMetadata(),
              'module.Fiery' => Fiery::getExceptionMetadata(),
              'module.Serious' => Serious::getExceptionMetadata(),
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

