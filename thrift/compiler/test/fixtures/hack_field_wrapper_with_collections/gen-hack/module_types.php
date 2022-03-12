<?hh
/**
 * Autogenerated by Thrift
 *
 * DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING
 *  @generated
 */

/**
 * Original thrift struct:-
 * AnnotationStruct
 */
class AnnotationStruct implements \IThriftSyncStruct, \IThriftShapishSyncStruct {
  use \ThriftSerializationTrait;

  const dict<int, this::TFieldSpec> SPEC = dict[
  ];
  const dict<string, int> FIELDMAP = dict[
  ];

  const type TConstructorShape = shape(
  );

  const type TShape = shape(
  );
  const int STRUCTURAL_ID = 957977401221134810;

  public function __construct(  )[] {
  }

  public static function withDefaultValues()[]: this {
    return new static();
  }

  public static function fromShape(self::TConstructorShape $shape)[]: this {
    return new static(
    );
  }

  public function getName()[]: string {
    return 'AnnotationStruct';
  }

  public static function getStructMetadata()[]: \tmeta_ThriftStruct {
    return tmeta_ThriftStruct::fromShape(
      shape(
        "name" => "module.AnnotationStruct",
        "is_union" => false,
      )
    );
  }

  public static function getAllStructuredAnnotations()[]: \TStructAnnotations {
    return shape(
      'struct' => dict[
        'facebook_thrift_annotation_Transitive' => facebook_thrift_annotation_Transitive::fromShape(
          shape(
          )
        ),
        '\facebook_thrift_annotation\FieldWrapper' => \facebook_thrift_annotation\FieldWrapper::fromShape(
          shape(
            "name" => "\MyFieldWrapper",
          )
        ),
      ],
      'fields' => dict[
      ],
    );
  }

  public static function __stringifyMapKeys<T>(Map<arraykey, T> $m)[]: Map<string, T> {
    $new = dict[];
    foreach ($m as $k => $v) {
      $new[(string)$k] = $v;
    }
    return new Map($new);
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

/**
 * Original thrift struct:-
 * MyStruct
 */
class MyStruct implements \IThriftSyncStruct, \IThriftShapishAsyncStruct {
  use \ThriftSerializationTrait;

  const dict<int, this::TFieldSpec> SPEC = dict[
    1 => shape(
      'var' => 'nested_struct',
      'type' => \TType::STRUCT,
      'class' => MyNestedStruct::class,
    ),
  ];
  const dict<string, int> FIELDMAP = dict[
    'nested_struct' => 1,
  ];

  const type TConstructorShape = shape(
    ?'nested_struct' => ?MyNestedStruct,
  );

  const type TShape = shape(
    ?'nested_struct' => ?MyNestedStruct::TShape,
  );
  const int STRUCTURAL_ID = 6466034702854646588;
  /**
   * Original thrift field:-
   * 1: struct module.MyNestedStruct nested_struct
   */
  public ?MyNestedStruct $nested_struct;

  public function __construct(?MyNestedStruct $nested_struct = null  )[] {
    $this->nested_struct = $nested_struct;
  }

  public static function withDefaultValues()[]: this {
    return new static();
  }

  public static function fromShape(self::TConstructorShape $shape)[]: this {
    return new static(
      Shapes::idx($shape, 'nested_struct'),
    );
  }

  public function getName()[]: string {
    return 'MyStruct';
  }

  public static function getStructMetadata()[]: \tmeta_ThriftStruct {
    return tmeta_ThriftStruct::fromShape(
      shape(
        "name" => "module.MyStruct",
        "fields" => vec[
          tmeta_ThriftField::fromShape(
            shape(
              "id" => 1,
              "type" => tmeta_ThriftType::fromShape(
                shape(
                  "t_typedef" => tmeta_ThriftTypedefType::fromShape(
                    shape(
                      "name" => "module.MyNestedStruct",
                      "underlyingType" => tmeta_ThriftType::fromShape(
                        shape(
                          "t_struct" => tmeta_ThriftStructType::fromShape(
                            shape(
                              "name" => "module.MyNestedStruct",
                            )
                          ),
                        )
                      ),
                    )
                  ),
                )
              ),
              "name" => "nested_struct",
            )
          ),
        ],
        "is_union" => false,
      )
    );
  }

  public static function getAllStructuredAnnotations()[]: \TStructAnnotations {
    return shape(
      'struct' => dict[],
      'fields' => dict[
      ],
    );
  }

  public static function __stringifyMapKeys<T>(Map<arraykey, T> $m)[]: Map<string, T> {
    $new = dict[];
    foreach ($m as $k => $v) {
      $new[(string)$k] = $v;
    }
    return new Map($new);
  }

  public static function __fromShape(self::TShape $shape)[]: this {
    return new static(
      Shapes::idx($shape, 'nested_struct') === null ? null : (MyNestedStruct::__fromShape($shape['nested_struct'])),
    );
  }

  public function __toShape()[]: self::TShape {
    return shape(
      'nested_struct' => $this->nested_struct?->__toShape(),
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

    if (idx($parsed, 'nested_struct') !== null) {
      $_tmp0 = json_encode(/* HH_FIXME[4110] */ $parsed['nested_struct']);
      $_tmp1 = MyNestedStruct::withDefaultValues();
      $_tmp1->readFromJson($_tmp0);
      $this->nested_struct = $_tmp1;
    }    
  }

}

/**
 * Original thrift struct:-
 * MyNestedStruct
 */
class MyNestedStruct implements \IThriftAsyncStruct, \IThriftShapishAsyncStruct {
  use \ThriftSerializationTrait;

  const dict<int, this::TFieldSpec> SPEC = dict[
    1 => shape(
      'var' => 'wrapped_field',
      'is_wrapped' => true,
      'type' => \TType::I64,
    ),
    2 => shape(
      'var' => 'annotated_field',
      'is_wrapped' => true,
      'type' => \TType::I64,
    ),
    3 => shape(
      'var' => 'adapted_type',
      'adapter' => \MyAdapter::class,
      'type' => \TType::I64,
    ),
  ];
  const dict<string, int> FIELDMAP = dict[
    'wrapped_field' => 1,
    'annotated_field' => 2,
    'adapted_type' => 3,
  ];

  const type TConstructorShape = shape(
    ?'wrapped_field' => ?int,
    ?'annotated_field' => ?int,
    ?'adapted_type' => ?\MyAdapter::THackType,
  );

  const type TShape = shape(
    'wrapped_field' => int,
    'annotated_field' => int,
    'adapted_type' => \MyAdapter::THackType,
  );
  const int STRUCTURAL_ID = 3672151039112827748;
  /**
   * Original thrift field:-
   * 1: i64 wrapped_field
   */
  private ?\MyFieldWrapper<int, this> $wrapped_field;

  public function get_wrapped_field()[]: \MyFieldWrapper<int, this> {
    return $this->wrapped_field as nonnull;
  }

  /**
   * Original thrift field:-
   * 2: i64 annotated_field
   */
  private ?\MyFieldWrapper<int, this> $annotated_field;

  public function get_annotated_field()[]: \MyFieldWrapper<int, this> {
    return $this->annotated_field as nonnull;
  }

  /**
   * Original thrift field:-
   * 3: i64 adapted_type
   */
  public \MyAdapter::THackType $adapted_type;

  public function __construct()[] {
    $this->adapted_type = \MyAdapter::fromThrift(0);
    $this->wrapped_field = \MyFieldWrapper::fromThrift_DO_NOT_USE_THRIFT_INTERNAL<int, this>(0, 1, $this);
    $this->annotated_field = \MyFieldWrapper::fromThrift_DO_NOT_USE_THRIFT_INTERNAL<int, this>(0, 2, $this);
  }

  public static function withDefaultValues()[]: this {
    return new static();
  }

  public static function fromShape(self::TConstructorShape $shape)[]: this {
    return new static(
      Shapes::idx($shape, 'wrapped_field'),
      Shapes::idx($shape, 'annotated_field'),
      Shapes::idx($shape, 'adapted_type'),
    );
  }

  public function getName()[]: string {
    return 'MyNestedStruct';
  }

  public static function getStructMetadata()[]: \tmeta_ThriftStruct {
    return tmeta_ThriftStruct::fromShape(
      shape(
        "name" => "module.MyNestedStruct",
        "fields" => vec[
          tmeta_ThriftField::fromShape(
            shape(
              "id" => 1,
              "type" => tmeta_ThriftType::fromShape(
                shape(
                  "t_primitive" => tmeta_ThriftPrimitiveType::THRIFT_I64_TYPE,
                )
              ),
              "name" => "wrapped_field",
            )
          ),
          tmeta_ThriftField::fromShape(
            shape(
              "id" => 2,
              "type" => tmeta_ThriftType::fromShape(
                shape(
                  "t_primitive" => tmeta_ThriftPrimitiveType::THRIFT_I64_TYPE,
                )
              ),
              "name" => "annotated_field",
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
              "name" => "adapted_type",
            )
          ),
        ],
        "is_union" => false,
      )
    );
  }

  public static function getAllStructuredAnnotations()[]: \TStructAnnotations {
    return shape(
      'struct' => dict[],
      'fields' => dict[
        'wrapped_field' => shape(
          'field' => dict[
            '\facebook_thrift_annotation\FieldWrapper' => \facebook_thrift_annotation\FieldWrapper::fromShape(
              shape(
                "name" => "\MyFieldWrapper",
              )
            ),
          ],
          'type' => dict[],
        ),
        'annotated_field' => shape(
          'field' => dict[
            'AnnotationStruct' => AnnotationStruct::fromShape(
              shape(
              )
            ),
          ],
          'type' => dict[],
        ),
      ],
    );
  }

  public static function __stringifyMapKeys<T>(Map<arraykey, T> $m)[]: Map<string, T> {
    $new = dict[];
    foreach ($m as $k => $v) {
      $new[(string)$k] = $v;
    }
    return new Map($new);
  }

  public static function __fromShape(self::TShape $shape)[]: this {
    return new static(
      $shape['wrapped_field'],
      $shape['annotated_field'],
      $shape['adapted_type'],
    );
  }

  public function __toShape()[]: self::TShape {
    return shape(
      'wrapped_field' => $this->wrapped_field,
      'annotated_field' => $this->annotated_field,
      'adapted_type' => $this->adapted_type,
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

    if (idx($parsed, 'wrapped_field') !== null) {
      $this->wrapped_field = /* HH_FIXME[4110] */ $parsed['wrapped_field'];
    }    
    if (idx($parsed, 'annotated_field') !== null) {
      $this->annotated_field = /* HH_FIXME[4110] */ $parsed['annotated_field'];
    }    
    if (idx($parsed, 'adapted_type') !== null) {
      $this->adapted_type = /* HH_FIXME[4110] */ $parsed['adapted_type'];
    }    
  }

  private static function __hackAdapterTypeChecks()[]: void {
    \ThriftUtil::requireSameType<\MyAdapter::TThriftType, int>();
  }

}

/**
 * Original thrift struct:-
 * MyComplexStruct
 */
class MyComplexStruct implements \IThriftSyncStruct, \IThriftShapishAsyncStruct {
  use \ThriftSerializationTrait;

  const dict<int, this::TFieldSpec> SPEC = dict[
    1 => shape(
      'var' => 'map_of_string_to_MyStruct',
      'type' => \TType::MAP,
      'ktype' => \TType::STRING,
      'vtype' => \TType::STRUCT,
      'key' => shape(
        'type' => \TType::STRING,
      ),
      'val' => shape(
        'type' => \TType::STRUCT,
        'class' => MyStruct::class,
      ),
      'format' => 'collection',
    ),
    2 => shape(
      'var' => 'map_of_string_to_list_of_MyStruct',
      'type' => \TType::MAP,
      'ktype' => \TType::STRING,
      'vtype' => \TType::LST,
      'key' => shape(
        'type' => \TType::STRING,
      ),
      'val' => shape(
        'type' => \TType::LST,
        'etype' => \TType::STRUCT,
        'elem' => shape(
          'type' => \TType::STRUCT,
          'class' => MyStruct::class,
        ),
        'format' => 'collection',
      ),
      'format' => 'collection',
    ),
    3 => shape(
      'var' => 'map_of_string_to_map_of_string_to_i32',
      'type' => \TType::MAP,
      'ktype' => \TType::STRING,
      'vtype' => \TType::MAP,
      'key' => shape(
        'type' => \TType::STRING,
      ),
      'val' => shape(
        'type' => \TType::MAP,
        'ktype' => \TType::STRING,
        'vtype' => \TType::I32,
        'key' => shape(
          'type' => \TType::STRING,
        ),
        'val' => shape(
          'type' => \TType::I32,
        ),
        'format' => 'collection',
      ),
      'format' => 'collection',
    ),
    4 => shape(
      'var' => 'map_of_string_to_map_of_string_to_MyStruct',
      'type' => \TType::MAP,
      'ktype' => \TType::STRING,
      'vtype' => \TType::MAP,
      'key' => shape(
        'type' => \TType::STRING,
      ),
      'val' => shape(
        'type' => \TType::MAP,
        'ktype' => \TType::STRING,
        'vtype' => \TType::STRUCT,
        'key' => shape(
          'type' => \TType::STRING,
        ),
        'val' => shape(
          'type' => \TType::STRUCT,
          'class' => MyStruct::class,
        ),
        'format' => 'collection',
      ),
      'format' => 'collection',
    ),
    5 => shape(
      'var' => 'list_of_map_of_string_to_list_of_MyStruct',
      'type' => \TType::LST,
      'etype' => \TType::MAP,
      'elem' => shape(
        'type' => \TType::MAP,
        'ktype' => \TType::STRING,
        'vtype' => \TType::LST,
        'key' => shape(
          'type' => \TType::STRING,
        ),
        'val' => shape(
          'type' => \TType::LST,
          'etype' => \TType::STRUCT,
          'elem' => shape(
            'type' => \TType::STRUCT,
            'class' => MyStruct::class,
          ),
          'format' => 'collection',
        ),
        'format' => 'collection',
      ),
      'format' => 'collection',
    ),
    6 => shape(
      'var' => 'list_of_map_of_string_to_MyStruct',
      'type' => \TType::LST,
      'etype' => \TType::MAP,
      'elem' => shape(
        'type' => \TType::MAP,
        'ktype' => \TType::STRING,
        'vtype' => \TType::STRUCT,
        'key' => shape(
          'type' => \TType::STRING,
        ),
        'val' => shape(
          'type' => \TType::STRUCT,
          'class' => MyStruct::class,
        ),
        'format' => 'collection',
      ),
      'format' => 'collection',
    ),
  ];
  const dict<string, int> FIELDMAP = dict[
    'map_of_string_to_MyStruct' => 1,
    'map_of_string_to_list_of_MyStruct' => 2,
    'map_of_string_to_map_of_string_to_i32' => 3,
    'map_of_string_to_map_of_string_to_MyStruct' => 4,
    'list_of_map_of_string_to_list_of_MyStruct' => 5,
    'list_of_map_of_string_to_MyStruct' => 6,
  ];

  const type TConstructorShape = shape(
    ?'map_of_string_to_MyStruct' => ?Map<string, MyStruct>,
    ?'map_of_string_to_list_of_MyStruct' => ?Map<string, Vector<MyStruct>>,
    ?'map_of_string_to_map_of_string_to_i32' => ?Map<string, Map<string, int>>,
    ?'map_of_string_to_map_of_string_to_MyStruct' => ?Map<string, Map<string, MyStruct>>,
    ?'list_of_map_of_string_to_list_of_MyStruct' => ?Vector<Map<string, Vector<MyStruct>>>,
    ?'list_of_map_of_string_to_MyStruct' => ?Vector<Map<string, MyStruct>>,
  );

  const type TShape = shape(
    'map_of_string_to_MyStruct' => dict<arraykey, MyStruct::TShape>,
    'map_of_string_to_list_of_MyStruct' => dict<arraykey, vec<MyStruct::TShape>>,
    'map_of_string_to_map_of_string_to_i32' => dict<arraykey, dict<arraykey, int>>,
    'map_of_string_to_map_of_string_to_MyStruct' => dict<arraykey, dict<arraykey, MyStruct::TShape>>,
    'list_of_map_of_string_to_list_of_MyStruct' => vec<dict<arraykey, vec<MyStruct::TShape>>>,
    'list_of_map_of_string_to_MyStruct' => vec<dict<arraykey, MyStruct::TShape>>,
  );
  const int STRUCTURAL_ID = 8703607584527805207;
  /**
   * Original thrift field:-
   * 1: map<string, struct module.MyStruct> map_of_string_to_MyStruct
   */
  public Map<string, MyStruct> $map_of_string_to_MyStruct;
  /**
   * Original thrift field:-
   * 2: map<string, list<struct module.MyStruct>> map_of_string_to_list_of_MyStruct
   */
  public Map<string, Vector<MyStruct>> $map_of_string_to_list_of_MyStruct;
  /**
   * Original thrift field:-
   * 3: map<string, map<string, i32>> map_of_string_to_map_of_string_to_i32
   */
  public Map<string, Map<string, int>> $map_of_string_to_map_of_string_to_i32;
  /**
   * Original thrift field:-
   * 4: map<string, map<string, struct module.MyStruct>> map_of_string_to_map_of_string_to_MyStruct
   */
  public Map<string, Map<string, MyStruct>> $map_of_string_to_map_of_string_to_MyStruct;
  /**
   * Original thrift field:-
   * 5: list<map<string, list<struct module.MyStruct>>> list_of_map_of_string_to_list_of_MyStruct
   */
  public Vector<Map<string, Vector<MyStruct>>> $list_of_map_of_string_to_list_of_MyStruct;
  /**
   * Original thrift field:-
   * 6: list<map<string, struct module.MyStruct>> list_of_map_of_string_to_MyStruct
   */
  public Vector<Map<string, MyStruct>> $list_of_map_of_string_to_MyStruct;

  public function __construct(?Map<string, MyStruct> $map_of_string_to_MyStruct = null, ?Map<string, Vector<MyStruct>> $map_of_string_to_list_of_MyStruct = null, ?Map<string, Map<string, int>> $map_of_string_to_map_of_string_to_i32 = null, ?Map<string, Map<string, MyStruct>> $map_of_string_to_map_of_string_to_MyStruct = null, ?Vector<Map<string, Vector<MyStruct>>> $list_of_map_of_string_to_list_of_MyStruct = null, ?Vector<Map<string, MyStruct>> $list_of_map_of_string_to_MyStruct = null  )[] {
    $this->map_of_string_to_MyStruct = $map_of_string_to_MyStruct ?? Map {};
    $this->map_of_string_to_list_of_MyStruct = $map_of_string_to_list_of_MyStruct ?? Map {};
    $this->map_of_string_to_map_of_string_to_i32 = $map_of_string_to_map_of_string_to_i32 ?? Map {};
    $this->map_of_string_to_map_of_string_to_MyStruct = $map_of_string_to_map_of_string_to_MyStruct ?? Map {};
    $this->list_of_map_of_string_to_list_of_MyStruct = $list_of_map_of_string_to_list_of_MyStruct ?? Vector {};
    $this->list_of_map_of_string_to_MyStruct = $list_of_map_of_string_to_MyStruct ?? Vector {};
  }

  public static function withDefaultValues()[]: this {
    return new static();
  }

  public static function fromShape(self::TConstructorShape $shape)[]: this {
    return new static(
      Shapes::idx($shape, 'map_of_string_to_MyStruct'),
      Shapes::idx($shape, 'map_of_string_to_list_of_MyStruct'),
      Shapes::idx($shape, 'map_of_string_to_map_of_string_to_i32'),
      Shapes::idx($shape, 'map_of_string_to_map_of_string_to_MyStruct'),
      Shapes::idx($shape, 'list_of_map_of_string_to_list_of_MyStruct'),
      Shapes::idx($shape, 'list_of_map_of_string_to_MyStruct'),
    );
  }

  public function getName()[]: string {
    return 'MyComplexStruct';
  }

  public static function getStructMetadata()[]: \tmeta_ThriftStruct {
    return tmeta_ThriftStruct::fromShape(
      shape(
        "name" => "module.MyComplexStruct",
        "fields" => vec[
          tmeta_ThriftField::fromShape(
            shape(
              "id" => 1,
              "type" => tmeta_ThriftType::fromShape(
                shape(
                  "t_map" => tmeta_ThriftMapType::fromShape(
                    shape(
                      "keyType" => tmeta_ThriftType::fromShape(
                        shape(
                          "t_primitive" => tmeta_ThriftPrimitiveType::THRIFT_STRING_TYPE,
                        )
                      ),
                      "valueType" => tmeta_ThriftType::fromShape(
                        shape(
                          "t_struct" => tmeta_ThriftStructType::fromShape(
                            shape(
                              "name" => "module.MyStruct",
                            )
                          ),
                        )
                      ),
                    )
                  ),
                )
              ),
              "name" => "map_of_string_to_MyStruct",
            )
          ),
          tmeta_ThriftField::fromShape(
            shape(
              "id" => 2,
              "type" => tmeta_ThriftType::fromShape(
                shape(
                  "t_map" => tmeta_ThriftMapType::fromShape(
                    shape(
                      "keyType" => tmeta_ThriftType::fromShape(
                        shape(
                          "t_primitive" => tmeta_ThriftPrimitiveType::THRIFT_STRING_TYPE,
                        )
                      ),
                      "valueType" => tmeta_ThriftType::fromShape(
                        shape(
                          "t_list" => tmeta_ThriftListType::fromShape(
                            shape(
                              "valueType" => tmeta_ThriftType::fromShape(
                                shape(
                                  "t_struct" => tmeta_ThriftStructType::fromShape(
                                    shape(
                                      "name" => "module.MyStruct",
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
              "name" => "map_of_string_to_list_of_MyStruct",
            )
          ),
          tmeta_ThriftField::fromShape(
            shape(
              "id" => 3,
              "type" => tmeta_ThriftType::fromShape(
                shape(
                  "t_map" => tmeta_ThriftMapType::fromShape(
                    shape(
                      "keyType" => tmeta_ThriftType::fromShape(
                        shape(
                          "t_primitive" => tmeta_ThriftPrimitiveType::THRIFT_STRING_TYPE,
                        )
                      ),
                      "valueType" => tmeta_ThriftType::fromShape(
                        shape(
                          "t_map" => tmeta_ThriftMapType::fromShape(
                            shape(
                              "keyType" => tmeta_ThriftType::fromShape(
                                shape(
                                  "t_primitive" => tmeta_ThriftPrimitiveType::THRIFT_STRING_TYPE,
                                )
                              ),
                              "valueType" => tmeta_ThriftType::fromShape(
                                shape(
                                  "t_primitive" => tmeta_ThriftPrimitiveType::THRIFT_I32_TYPE,
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
              "name" => "map_of_string_to_map_of_string_to_i32",
            )
          ),
          tmeta_ThriftField::fromShape(
            shape(
              "id" => 4,
              "type" => tmeta_ThriftType::fromShape(
                shape(
                  "t_map" => tmeta_ThriftMapType::fromShape(
                    shape(
                      "keyType" => tmeta_ThriftType::fromShape(
                        shape(
                          "t_primitive" => tmeta_ThriftPrimitiveType::THRIFT_STRING_TYPE,
                        )
                      ),
                      "valueType" => tmeta_ThriftType::fromShape(
                        shape(
                          "t_map" => tmeta_ThriftMapType::fromShape(
                            shape(
                              "keyType" => tmeta_ThriftType::fromShape(
                                shape(
                                  "t_primitive" => tmeta_ThriftPrimitiveType::THRIFT_STRING_TYPE,
                                )
                              ),
                              "valueType" => tmeta_ThriftType::fromShape(
                                shape(
                                  "t_struct" => tmeta_ThriftStructType::fromShape(
                                    shape(
                                      "name" => "module.MyStruct",
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
              "name" => "map_of_string_to_map_of_string_to_MyStruct",
            )
          ),
          tmeta_ThriftField::fromShape(
            shape(
              "id" => 5,
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
                                  "t_primitive" => tmeta_ThriftPrimitiveType::THRIFT_STRING_TYPE,
                                )
                              ),
                              "valueType" => tmeta_ThriftType::fromShape(
                                shape(
                                  "t_list" => tmeta_ThriftListType::fromShape(
                                    shape(
                                      "valueType" => tmeta_ThriftType::fromShape(
                                        shape(
                                          "t_struct" => tmeta_ThriftStructType::fromShape(
                                            shape(
                                              "name" => "module.MyStruct",
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
                )
              ),
              "name" => "list_of_map_of_string_to_list_of_MyStruct",
            )
          ),
          tmeta_ThriftField::fromShape(
            shape(
              "id" => 6,
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
                                  "t_primitive" => tmeta_ThriftPrimitiveType::THRIFT_STRING_TYPE,
                                )
                              ),
                              "valueType" => tmeta_ThriftType::fromShape(
                                shape(
                                  "t_struct" => tmeta_ThriftStructType::fromShape(
                                    shape(
                                      "name" => "module.MyStruct",
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
              "name" => "list_of_map_of_string_to_MyStruct",
            )
          ),
        ],
        "is_union" => false,
      )
    );
  }

  public static function getAllStructuredAnnotations()[]: \TStructAnnotations {
    return shape(
      'struct' => dict[],
      'fields' => dict[
      ],
    );
  }

  public static function __stringifyMapKeys<T>(Map<arraykey, T> $m)[]: Map<string, T> {
    $new = dict[];
    foreach ($m as $k => $v) {
      $new[(string)$k] = $v;
    }
    return new Map($new);
  }

  public static function __fromShape(self::TShape $shape)[]: this {
    return new static(
      self::__stringifyMapKeys(new Map($shape['map_of_string_to_MyStruct']))->map(
        $val0 ==> MyStruct::__fromShape($val0),
      ),
      self::__stringifyMapKeys(new Map($shape['map_of_string_to_list_of_MyStruct']))->map(
        $val1 ==> (new Vector($val1))->map(
          $val2 ==> MyStruct::__fromShape($val2),
        ),
      ),
      self::__stringifyMapKeys(new Map($shape['map_of_string_to_map_of_string_to_i32']))->map(
        $val3 ==> self::__stringifyMapKeys(new Map($val3)),
      ),
      self::__stringifyMapKeys(new Map($shape['map_of_string_to_map_of_string_to_MyStruct']))->map(
        $val4 ==> self::__stringifyMapKeys(new Map($val4))->map(
          $val5 ==> MyStruct::__fromShape($val5),
        ),
      ),
      (new Vector($shape['list_of_map_of_string_to_list_of_MyStruct']))->map(
        $val6 ==> self::__stringifyMapKeys(new Map($val6))->map(
          $val7 ==> (new Vector($val7))->map(
            $val8 ==> MyStruct::__fromShape($val8),
          ),
        ),
      ),
      (new Vector($shape['list_of_map_of_string_to_MyStruct']))->map(
        $val9 ==> self::__stringifyMapKeys(new Map($val9))->map(
          $val10 ==> MyStruct::__fromShape($val10),
        ),
      ),
    );
  }

  public function __toShape()[]: self::TShape {
    return shape(
      'map_of_string_to_MyStruct' => $this->map_of_string_to_MyStruct->map(
        ($_val0) ==> $_val0->__toShape(),
      )
        |> dict($$),
      'map_of_string_to_list_of_MyStruct' => $this->map_of_string_to_list_of_MyStruct->map(
        ($_val0) ==> $_val0->map(
          ($_val1) ==> $_val1->__toShape(),
        )
          |> vec($$),
      )
        |> dict($$),
      'map_of_string_to_map_of_string_to_i32' => $this->map_of_string_to_map_of_string_to_i32->map(
        ($_val0) ==> dict($_val0),
      )
        |> dict($$),
      'map_of_string_to_map_of_string_to_MyStruct' => $this->map_of_string_to_map_of_string_to_MyStruct->map(
        ($_val0) ==> $_val0->map(
          ($_val1) ==> $_val1->__toShape(),
        )
          |> dict($$),
      )
        |> dict($$),
      'list_of_map_of_string_to_list_of_MyStruct' => $this->list_of_map_of_string_to_list_of_MyStruct->map(
        ($_val0) ==> $_val0->map(
          ($_val1) ==> $_val1->map(
            ($_val2) ==> $_val2->__toShape(),
          )
            |> vec($$),
        )
          |> dict($$),
      )
        |> vec($$),
      'list_of_map_of_string_to_MyStruct' => $this->list_of_map_of_string_to_MyStruct->map(
        ($_val0) ==> $_val0->map(
          ($_val1) ==> $_val1->__toShape(),
        )
          |> dict($$),
      )
        |> vec($$),
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

    if (idx($parsed, 'map_of_string_to_MyStruct') !== null) {
      $_json3 = /* HH_FIXME[4110] */ $parsed['map_of_string_to_MyStruct'];
      $_container4 = Map {};
      foreach(/* HH_FIXME[4110] */ $_json3 as $_key1 => $_value2) {
        $_value5 = MyStruct::withDefaultValues();
        $_tmp6 = json_encode($_value2);
        $_tmp7 = MyStruct::withDefaultValues();
        $_tmp7->readFromJson($_tmp6);
        $_value5 = $_tmp7;
        $_container4[$_key1] = $_value5;
      }
      $this->map_of_string_to_MyStruct = $_container4;
    }    
    if (idx($parsed, 'map_of_string_to_list_of_MyStruct') !== null) {
      $_json11 = /* HH_FIXME[4110] */ $parsed['map_of_string_to_list_of_MyStruct'];
      $_container12 = Map {};
      foreach(/* HH_FIXME[4110] */ $_json11 as $_key9 => $_value10) {
        $_value13 = Vector {};
        $_json17 = $_value10;
        $_container18 = Vector {};
        foreach(/* HH_FIXME[4110] */ $_json17 as $_key15 => $_value16) {
          $_elem19 = MyStruct::withDefaultValues();
          $_tmp20 = json_encode($_value16);
          $_tmp21 = MyStruct::withDefaultValues();
          $_tmp21->readFromJson($_tmp20);
          $_elem19 = $_tmp21;
          $_container18 []= $_elem19;
        }
        $_value13 = $_container18;
        $_container12[$_key9] = $_value13;
      }
      $this->map_of_string_to_list_of_MyStruct = $_container12;
    }    
    if (idx($parsed, 'map_of_string_to_map_of_string_to_i32') !== null) {
      $_json25 = /* HH_FIXME[4110] */ $parsed['map_of_string_to_map_of_string_to_i32'];
      $_container26 = Map {};
      foreach(/* HH_FIXME[4110] */ $_json25 as $_key23 => $_value24) {
        $_value27 = Map {};
        $_json31 = $_value24;
        $_container32 = Map {};
        foreach(/* HH_FIXME[4110] */ $_json31 as $_key29 => $_value30) {
          $_value33 = 0;
          $_tmp34 = (int)$_value30;
          if ($_tmp34 > 0x7fffffff) {
            throw new \TProtocolException("number exceeds limit in field");
          } else {
            $_value33 = (int)$_tmp34;
          }
          $_container32[$_key29] = $_value33;
        }
        $_value27 = $_container32;
        $_container26[$_key23] = $_value27;
      }
      $this->map_of_string_to_map_of_string_to_i32 = $_container26;
    }    
    if (idx($parsed, 'map_of_string_to_map_of_string_to_MyStruct') !== null) {
      $_json38 = /* HH_FIXME[4110] */ $parsed['map_of_string_to_map_of_string_to_MyStruct'];
      $_container39 = Map {};
      foreach(/* HH_FIXME[4110] */ $_json38 as $_key36 => $_value37) {
        $_value40 = Map {};
        $_json44 = $_value37;
        $_container45 = Map {};
        foreach(/* HH_FIXME[4110] */ $_json44 as $_key42 => $_value43) {
          $_value46 = MyStruct::withDefaultValues();
          $_tmp47 = json_encode($_value43);
          $_tmp48 = MyStruct::withDefaultValues();
          $_tmp48->readFromJson($_tmp47);
          $_value46 = $_tmp48;
          $_container45[$_key42] = $_value46;
        }
        $_value40 = $_container45;
        $_container39[$_key36] = $_value40;
      }
      $this->map_of_string_to_map_of_string_to_MyStruct = $_container39;
    }    
    if (idx($parsed, 'list_of_map_of_string_to_list_of_MyStruct') !== null) {
      $_json52 = /* HH_FIXME[4110] */ $parsed['list_of_map_of_string_to_list_of_MyStruct'];
      $_container53 = Vector {};
      foreach(/* HH_FIXME[4110] */ $_json52 as $_key50 => $_value51) {
        $_elem54 = Map {};
        $_json58 = $_value51;
        $_container59 = Map {};
        foreach(/* HH_FIXME[4110] */ $_json58 as $_key56 => $_value57) {
          $_value60 = Vector {};
          $_json64 = $_value57;
          $_container65 = Vector {};
          foreach(/* HH_FIXME[4110] */ $_json64 as $_key62 => $_value63) {
            $_elem66 = MyStruct::withDefaultValues();
            $_tmp67 = json_encode($_value63);
            $_tmp68 = MyStruct::withDefaultValues();
            $_tmp68->readFromJson($_tmp67);
            $_elem66 = $_tmp68;
            $_container65 []= $_elem66;
          }
          $_value60 = $_container65;
          $_container59[$_key56] = $_value60;
        }
        $_elem54 = $_container59;
        $_container53 []= $_elem54;
      }
      $this->list_of_map_of_string_to_list_of_MyStruct = $_container53;
    }    
    if (idx($parsed, 'list_of_map_of_string_to_MyStruct') !== null) {
      $_json72 = /* HH_FIXME[4110] */ $parsed['list_of_map_of_string_to_MyStruct'];
      $_container73 = Vector {};
      foreach(/* HH_FIXME[4110] */ $_json72 as $_key70 => $_value71) {
        $_elem74 = Map {};
        $_json78 = $_value71;
        $_container79 = Map {};
        foreach(/* HH_FIXME[4110] */ $_json78 as $_key76 => $_value77) {
          $_value80 = MyStruct::withDefaultValues();
          $_tmp81 = json_encode($_value77);
          $_tmp82 = MyStruct::withDefaultValues();
          $_tmp82->readFromJson($_tmp81);
          $_value80 = $_tmp82;
          $_container79[$_key76] = $_value80;
        }
        $_elem74 = $_container79;
        $_container73 []= $_elem74;
      }
      $this->list_of_map_of_string_to_MyStruct = $_container73;
    }    
  }

}

