<?hh
/**
 * Autogenerated by Thrift
 *
 * DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING
 *  @generated
 */

/**
 * Original thrift service:-
 * InteractLocally
 */
interface InteractLocallyAsyncIf extends \IThriftAsyncIf {
}

/**
 * Original thrift service:-
 * InteractLocally
 */
interface InteractLocallyIf extends \IThriftSyncIf {
}

/**
 * Original thrift service:-
 * InteractLocally
 */
interface InteractLocallyAsyncClientIf extends InteractLocallyAsyncIf {
}

/**
 * Original thrift service:-
 * InteractLocally
 */
interface InteractLocallyClientIf extends \IThriftSyncIf {
}

/**
 * Original thrift service:-
 * InteractLocally
 */
trait InteractLocallyClientBase {
  require extends \ThriftClientBase;

  /* interaction handlers factory methods */
  public function createSharedInteraction(): InteractLocally_SharedInteraction {
    $interaction = new InteractLocally_SharedInteraction($this->input_, $this->output_, $this->channel_);
    $interaction->setAsyncHandler($this->asyncHandler_)->setEventHandler($this->eventHandler_);
    return $interaction;
  }

}

class InteractLocallyAsyncClient extends \ThriftClientBase implements InteractLocallyAsyncClientIf {
  use InteractLocallyClientBase;

}

class InteractLocallyClient extends \ThriftClientBase implements InteractLocallyClientIf {
  use InteractLocallyClientBase;

}

// INTERACTION HANDLERS

class InteractLocally_SharedInteraction extends \ThriftClientBase {
  private \InteractionId $interactionId;

  public function __construct(\TProtocol $input, ?\TProtocol $output = null, ?\IThriftMigrationAsyncChannel $channel = null)[leak_safe] {
    parent::__construct($input, $output, $channel);
    if ($this->channel_ is nonnull) {
      $this->interactionId = $this->channel_->createInteraction("SharedInteraction");
    } else {
      throw new \Exception("The channel must be nonnull to create interactions.");
    }
  }

  /**
   * Original thrift definition:-
   * i32
   *   init();
   */
  public async function init(): Awaitable<int> {
    $hh_frame_metadata = $this->getHHFrameMetadata();
    if ($hh_frame_metadata !== null) {
      \HH\set_frame_metadata($hh_frame_metadata);
    }
    $rpc_options = $this->getAndResetOptions() ?? new \RpcOptions();
    $rpc_options = $rpc_options->setInteractionId($this->interactionId);
    $args = InteractLocally_SharedInteraction_init_args::withDefaultValues();
    await $this->asyncHandler_->genBefore("InteractLocally", "SharedInteraction.init", $args);
    $currentseqid = $this->sendImpl_init();
    return await $this->genAwaitResponse(InteractLocally_SharedInteraction_init_result::class, "init", false, $currentseqid, $rpc_options);
  }

  protected function sendImpl_init(): int {
    $currentseqid = $this->getNextSequenceID();
    $args = InteractLocally_SharedInteraction_init_args::withDefaultValues();
    try {
      $this->eventHandler_->preSend('SharedInteraction.init', $args, $currentseqid, 'InteractLocally');
      if ($this->output_ is \TBinaryProtocolAccelerated)
      {
        \thrift_protocol_write_binary($this->output_, 'SharedInteraction.init', \TMessageType::CALL, $args, $currentseqid, $this->output_->isStrictWrite(), false);
      }
      else if ($this->output_ is \TCompactProtocolAccelerated)
      {
        \thrift_protocol_write_compact2($this->output_, 'SharedInteraction.init', \TMessageType::CALL, $args, $currentseqid, false, \TCompactProtocolBase::VERSION);
      }
      else
      {
        $this->output_->writeMessageBegin('SharedInteraction.init', \TMessageType::CALL, $currentseqid);
        $args->write($this->output_);
        $this->output_->writeMessageEnd();
        $this->output_->getTransport()->flush();
      }
    } catch (\THandlerShortCircuitException $ex) {
      switch ($ex->resultType) {
        case \THandlerShortCircuitException::R_EXPECTED_EX:
        case \THandlerShortCircuitException::R_UNEXPECTED_EX:
          $this->eventHandler_->sendError('SharedInteraction.init', $args, $currentseqid, $ex->result);
          throw $ex->result;
        case \THandlerShortCircuitException::R_SUCCESS:
        default:
          $this->eventHandler_->postSend('SharedInteraction.init', $args, $currentseqid);
          return $currentseqid;
      }
    } catch (\Exception $ex) {
      $this->eventHandler_->sendError('SharedInteraction.init', $args, $currentseqid, $ex);
      throw $ex;
    }
    $this->eventHandler_->postSend('SharedInteraction.init', $args, $currentseqid);
    return $currentseqid;
  }
  /**
   * Original thrift definition:-
   * DoSomethingResult
   *   do_something();
   */
  public async function do_something(): Awaitable<DoSomethingResult> {
    $hh_frame_metadata = $this->getHHFrameMetadata();
    if ($hh_frame_metadata !== null) {
      \HH\set_frame_metadata($hh_frame_metadata);
    }
    $rpc_options = $this->getAndResetOptions() ?? new \RpcOptions();
    $rpc_options = $rpc_options->setInteractionId($this->interactionId);
    $args = InteractLocally_SharedInteraction_do_something_args::withDefaultValues();
    await $this->asyncHandler_->genBefore("InteractLocally", "SharedInteraction.do_something", $args);
    $currentseqid = $this->sendImpl_do_something();
    return await $this->genAwaitResponse(InteractLocally_SharedInteraction_do_something_result::class, "do_something", false, $currentseqid, $rpc_options);
  }

  protected function sendImpl_do_something(): int {
    $currentseqid = $this->getNextSequenceID();
    $args = InteractLocally_SharedInteraction_do_something_args::withDefaultValues();
    try {
      $this->eventHandler_->preSend('SharedInteraction.do_something', $args, $currentseqid, 'InteractLocally');
      if ($this->output_ is \TBinaryProtocolAccelerated)
      {
        \thrift_protocol_write_binary($this->output_, 'SharedInteraction.do_something', \TMessageType::CALL, $args, $currentseqid, $this->output_->isStrictWrite(), false);
      }
      else if ($this->output_ is \TCompactProtocolAccelerated)
      {
        \thrift_protocol_write_compact2($this->output_, 'SharedInteraction.do_something', \TMessageType::CALL, $args, $currentseqid, false, \TCompactProtocolBase::VERSION);
      }
      else
      {
        $this->output_->writeMessageBegin('SharedInteraction.do_something', \TMessageType::CALL, $currentseqid);
        $args->write($this->output_);
        $this->output_->writeMessageEnd();
        $this->output_->getTransport()->flush();
      }
    } catch (\THandlerShortCircuitException $ex) {
      switch ($ex->resultType) {
        case \THandlerShortCircuitException::R_EXPECTED_EX:
        case \THandlerShortCircuitException::R_UNEXPECTED_EX:
          $this->eventHandler_->sendError('SharedInteraction.do_something', $args, $currentseqid, $ex->result);
          throw $ex->result;
        case \THandlerShortCircuitException::R_SUCCESS:
        default:
          $this->eventHandler_->postSend('SharedInteraction.do_something', $args, $currentseqid);
          return $currentseqid;
      }
    } catch (\Exception $ex) {
      $this->eventHandler_->sendError('SharedInteraction.do_something', $args, $currentseqid, $ex);
      throw $ex;
    }
    $this->eventHandler_->postSend('SharedInteraction.do_something', $args, $currentseqid);
    return $currentseqid;
  }
  /**
   * Original thrift definition:-
   * void
   *   tear_down();
   */
  public async function tear_down(): Awaitable<void> {
    $hh_frame_metadata = $this->getHHFrameMetadata();
    if ($hh_frame_metadata !== null) {
      \HH\set_frame_metadata($hh_frame_metadata);
    }
    $rpc_options = $this->getAndResetOptions() ?? new \RpcOptions();
    $rpc_options = $rpc_options->setInteractionId($this->interactionId);
    $args = InteractLocally_SharedInteraction_tear_down_args::withDefaultValues();
    await $this->asyncHandler_->genBefore("InteractLocally", "SharedInteraction.tear_down", $args);
    $currentseqid = $this->sendImpl_tear_down();
    await $this->genAwaitResponse(InteractLocally_SharedInteraction_tear_down_result::class, "tear_down", true, $currentseqid, $rpc_options);
  }

  protected function sendImpl_tear_down(): int {
    $currentseqid = $this->getNextSequenceID();
    $args = InteractLocally_SharedInteraction_tear_down_args::withDefaultValues();
    try {
      $this->eventHandler_->preSend('SharedInteraction.tear_down', $args, $currentseqid, 'InteractLocally');
      if ($this->output_ is \TBinaryProtocolAccelerated)
      {
        \thrift_protocol_write_binary($this->output_, 'SharedInteraction.tear_down', \TMessageType::CALL, $args, $currentseqid, $this->output_->isStrictWrite(), false);
      }
      else if ($this->output_ is \TCompactProtocolAccelerated)
      {
        \thrift_protocol_write_compact2($this->output_, 'SharedInteraction.tear_down', \TMessageType::CALL, $args, $currentseqid, false, \TCompactProtocolBase::VERSION);
      }
      else
      {
        $this->output_->writeMessageBegin('SharedInteraction.tear_down', \TMessageType::CALL, $currentseqid);
        $args->write($this->output_);
        $this->output_->writeMessageEnd();
        $this->output_->getTransport()->flush();
      }
    } catch (\THandlerShortCircuitException $ex) {
      switch ($ex->resultType) {
        case \THandlerShortCircuitException::R_EXPECTED_EX:
        case \THandlerShortCircuitException::R_UNEXPECTED_EX:
          $this->eventHandler_->sendError('SharedInteraction.tear_down', $args, $currentseqid, $ex->result);
          throw $ex->result;
        case \THandlerShortCircuitException::R_SUCCESS:
        default:
          $this->eventHandler_->postSend('SharedInteraction.tear_down', $args, $currentseqid);
          return $currentseqid;
      }
    } catch (\Exception $ex) {
      $this->eventHandler_->sendError('SharedInteraction.tear_down', $args, $currentseqid, $ex);
      throw $ex;
    }
    $this->eventHandler_->postSend('SharedInteraction.tear_down', $args, $currentseqid);
    return $currentseqid;
  }
}

// HELPER FUNCTIONS AND STRUCTURES

class InteractLocally_SharedInteraction_init_args implements \IThriftSyncStruct, \IThriftStructMetadata {
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
    return 'InteractLocally_SharedInteraction_init_args';
  }

  public static function getStructMetadata()[]: \tmeta_ThriftStruct {
    return tmeta_ThriftStruct::fromShape(
      shape(
        "name" => "shared.init_args",
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

class InteractLocally_SharedInteraction_init_result extends \ThriftSyncStructWithResult implements \IThriftStructMetadata {
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
    return 'InteractLocally_SharedInteraction_init_result';
  }

  public static function getStructMetadata()[]: \tmeta_ThriftStruct {
    return tmeta_ThriftStruct::fromShape(
      shape(
        "name" => "shared.InteractLocally_SharedInteraction_init_result",
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

}

class InteractLocally_SharedInteraction_do_something_args implements \IThriftSyncStruct, \IThriftStructMetadata {
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
    return 'InteractLocally_SharedInteraction_do_something_args';
  }

  public static function getStructMetadata()[]: \tmeta_ThriftStruct {
    return tmeta_ThriftStruct::fromShape(
      shape(
        "name" => "shared.do_something_args",
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

class InteractLocally_SharedInteraction_do_something_result extends \ThriftSyncStructWithResult implements \IThriftStructMetadata {
  use \ThriftSerializationTrait;

  const type TResult = DoSomethingResult;

  const \ThriftStructTypes::TSpec SPEC = dict[
    0 => shape(
      'var' => 'success',
      'type' => \TType::STRUCT,
      'class' => DoSomethingResult::class,
    ),
  ];
  const dict<string, int> FIELDMAP = dict[
    'success' => 0,
  ];

  const type TConstructorShape = shape(
    ?'success' => ?this::TResult,
  );

  const int STRUCTURAL_ID = 1665402448780378721;
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
    return 'InteractLocally_SharedInteraction_do_something_result';
  }

  public static function getStructMetadata()[]: \tmeta_ThriftStruct {
    return tmeta_ThriftStruct::fromShape(
      shape(
        "name" => "shared.InteractLocally_SharedInteraction_do_something_result",
        "fields" => vec[
          tmeta_ThriftField::fromShape(
            shape(
              "id" => 0,
              "type" => tmeta_ThriftType::fromShape(
                shape(
                  "t_struct" => tmeta_ThriftStructType::fromShape(
                    shape(
                      "name" => "shared.DoSomethingResult",
                    )
                  ),
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

class InteractLocally_SharedInteraction_tear_down_args implements \IThriftSyncStruct, \IThriftStructMetadata {
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
    return 'InteractLocally_SharedInteraction_tear_down_args';
  }

  public static function getStructMetadata()[]: \tmeta_ThriftStruct {
    return tmeta_ThriftStruct::fromShape(
      shape(
        "name" => "shared.tear_down_args",
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

class InteractLocally_SharedInteraction_tear_down_result extends \ThriftSyncStructWithoutResult implements \IThriftStructMetadata {
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
    return 'InteractLocally_SharedInteraction_tear_down_result';
  }

  public static function getStructMetadata()[]: \tmeta_ThriftStruct {
    return tmeta_ThriftStruct::fromShape(
      shape(
        "name" => "shared.InteractLocally_SharedInteraction_tear_down_result",
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

class InteractLocallyStaticMetadata implements \IThriftServiceStaticMetadata {
  public static function getServiceMetadata()[]: \tmeta_ThriftService {
    return tmeta_ThriftService::fromShape(
      shape(
        "name" => "shared.InteractLocally",
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
                'name' => 'shared',
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
      ],
    );
  }
}

