<?hh
/**
 * Autogenerated by Thrift
 *
 * DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING
 *  @generated
 */

/**
 * Original thrift enum:-
 * MyEnum
 */
enum MyEnum: int {
  MyValue1 = 0;
  MyValue2 = 1;
  DOMAIN = 2;
}

class MyEnum_TEnumStaticMetadata implements \IThriftEnumStaticMetadata {
  public static function getEnumMetadata()[]: \tmeta_ThriftEnum {
    return tmeta_ThriftEnum::fromShape(
      shape(
        "name" => "module.MyEnum",
        "elements" => dict[
          0 => "MyValue1",
          1 => "MyValue2",
          2 => "DOMAIN",
        ],
      )
    );
  }

  public static function getAllStructuredAnnotations()[write_props]: \TEnumAnnotations {
    return shape(
      'enum' => dict[],
      'constants' => dict[
      ],
    );
  }
}

/**
 * Original thrift struct:-
 * MyStructNestedAnnotation
 */
class MyStructNestedAnnotation implements \IThriftSyncStruct, \IThriftStructMetadata, \IThriftShapishSyncStruct {
  use \ThriftSerializationTrait;

  const \ThriftStructTypes::TSpec SPEC = dict[
    1 => shape(
      'var' => 'name',
      'type' => \TType::STRING,
    ),
  ];
  const dict<string, int> FIELDMAP = dict[
    'name' => 1,
  ];

  const type TConstructorShape = shape(
    ?'name' => ?string,
  );

  const type TShape = shape(
    'name' => string,
    ...
  );
  const int STRUCTURAL_ID = 2593878277785201336;
  /**
   * Original thrift field:-
   * 1: string name
   */
  public string $name;

  public function __construct(?string $name = null)[] {
    $this->name = $name ?? '';
  }

  public static function withDefaultValues()[]: this {
    return new static();
  }

  public static function fromShape(self::TConstructorShape $shape)[]: this {
    return new static(
      Shapes::idx($shape, 'name'),
    );
  }

  public function getName()[]: string {
    return 'MyStructNestedAnnotation';
  }

  public function clearTerseFields()[write_props]: void {
  }

  public static function getStructMetadata()[]: \tmeta_ThriftStruct {
    return tmeta_ThriftStruct::fromShape(
      shape(
        "name" => "module.MyStructNestedAnnotation",
        "fields" => vec[
          tmeta_ThriftField::fromShape(
            shape(
              "id" => 1,
              "type" => tmeta_ThriftType::fromShape(
                shape(
                  "t_primitive" => tmeta_ThriftPrimitiveType::THRIFT_STRING_TYPE,
                )
              ),
              "name" => "name",
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

  public static function __fromShape(self::TShape $shape)[]: this {
    return new static(
      $shape['name'],
    );
  }

  public function __toShape()[]: self::TShape {
    return shape(
      'name' => $this->name,
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

    if (idx($parsed, 'name') !== null) {
      $this->name = HH\FIXME\UNSAFE_CAST<mixed, string>($parsed['name']);
    }
  }

}

enum MyUnionEnum: int {
  _EMPTY_ = 0;
}

/**
 * Original thrift struct:-
 * MyUnion
 */
class MyUnion implements \IThriftSyncStruct, \IThriftStructMetadata, \IThriftUnion<MyUnionEnum>, \IThriftShapishSyncStruct {
  use \ThriftUnionSerializationTrait;

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
  protected MyUnionEnum $_type = MyUnionEnum::_EMPTY_;

  public function __construct()[] {
    $this->_type = MyUnionEnum::_EMPTY_;
  }

  public static function withDefaultValues()[]: this {
    return new static();
  }

  public static function fromShape(self::TConstructorShape $shape)[]: this {
    return new static(
    );
  }

  public function getName()[]: string {
    return 'MyUnion';
  }

  public function getType()[]: MyUnionEnum {
    return $this->_type;
  }

  public function reset()[write_props]: void {
    switch ($this->_type) {
      case MyUnionEnum::_EMPTY_:
        break;
    }
    $this->_type = MyUnionEnum::_EMPTY_;
  }

  public function clearTerseFields()[write_props]: void {
  }

  public static function getStructMetadata()[]: \tmeta_ThriftStruct {
    return tmeta_ThriftStruct::fromShape(
      shape(
        "name" => "module.MyUnion",
        "is_union" => true,
      )
    );
  }

  public static function getAllStructuredAnnotations()[write_props]: \TStructAnnotations {
    return shape(
      'struct' => dict[
        '\thrift\annotation\cpp\Adapter' => \thrift\annotation\cpp\Adapter::fromShape(
          shape(
            "name" => "::StaticCast",
          )
        ),
      ],
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
    $this->_type = MyUnionEnum::_EMPTY_;
    $parsed = json_decode($jsonText, true);

    if ($parsed === null || !($parsed is KeyedContainer<_, _>)) {
      throw new \TProtocolException("Cannot parse the given json string.");
    }

  }

}

/**
 * Original thrift exception:-
 * MyException
 */
class MyException extends \TException implements \IThriftSyncStruct, \IThriftExceptionMetadata {
  use \ThriftSerializationTrait;

  const \ThriftStructTypes::TSpec SPEC = dict[
  ];
  const dict<string, int> FIELDMAP = dict[
  ];

  const type TConstructorShape = shape(
  );

  const int STRUCTURAL_ID = 957977401221134810;

  public function __construct()[] {
    parent::__construct();
  }

  public static function withDefaultValues()[]: this {
    return new static();
  }

  public static function fromShape(self::TConstructorShape $shape)[]: this {
    return new static(
    );
  }

  public function getName()[]: string {
    return 'MyException';
  }

  public function clearTerseFields()[write_props]: void {
  }

  public static function getExceptionMetadata()[]: \tmeta_ThriftException {
    return tmeta_ThriftException::fromShape(
      shape(
        "name" => "module.MyException",
      )
    );
  }

  public static function getAllStructuredAnnotations()[write_props]: \TStructAnnotations {
    return shape(
      'struct' => dict[
        '\thrift\annotation\cpp\Adapter' => \thrift\annotation\cpp\Adapter::fromShape(
          shape(
            "name" => "::StaticCast",
          )
        ),
      ],
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

/**
 * Original thrift struct:-
 * MyStruct
 */
<<\ThriftTypeInfo(shape('uri' => 'facebook.com/thrift/compiler/test/fixtures/basic-annotations/src/module/MyStruct')),\SomeClass(\AnotherClass::class)>>
class MyStruct implements \IThriftSyncStruct, \IThriftStructMetadata, \IThriftShapishSyncStruct {
  use \ThriftSerializationTrait;

  const \ThriftStructTypes::TSpec SPEC = dict[
    2 => shape(
      'var' => 'major',
      'type' => \TType::I64,
    ),
    1 => shape(
      'var' => 'package',
      'type' => \TType::STRING,
    ),
    3 => shape(
      'var' => 'annotation_with_quote',
      'type' => \TType::STRING,
    ),
    4 => shape(
      'var' => 'class_',
      'type' => \TType::STRING,
    ),
    5 => shape(
      'var' => 'annotation_with_trailing_comma',
      'type' => \TType::STRING,
    ),
    6 => shape(
      'var' => 'empty_annotations',
      'type' => \TType::STRING,
    ),
    7 => shape(
      'var' => 'my_enum',
      'type' => \TType::I32,
      'enum' => MyEnum::class,
    ),
    8 => shape(
      'var' => 'cpp_type_annotation',
      'type' => \TType::LST,
      'etype' => \TType::STRING,
      'elem' => shape(
        'type' => \TType::STRING,
      ),
      'format' => 'collection',
    ),
    9 => shape(
      'var' => 'my_union',
      'type' => \TType::STRUCT,
      'class' => MyUnion::class,
    ),
    10 => shape(
      'var' => 'my_id',
      'type' => \TType::I16,
    ),
  ];
  const dict<string, int> FIELDMAP = dict[
    'major' => 2,
    'package' => 1,
    'annotation_with_quote' => 3,
    'class_' => 4,
    'annotation_with_trailing_comma' => 5,
    'empty_annotations' => 6,
    'my_enum' => 7,
    'cpp_type_annotation' => 8,
    'my_union' => 9,
    'my_id' => 10,
  ];

  const type TConstructorShape = shape(
    ?'major' => ?int,
    ?'package' => ?string,
    ?'annotation_with_quote' => ?string,
    ?'class_' => ?string,
    ?'annotation_with_trailing_comma' => ?string,
    ?'empty_annotations' => ?string,
    ?'my_enum' => ?MyEnum,
    ?'cpp_type_annotation' => ?Vector<string>,
    ?'my_union' => ?MyUnion,
    ?'my_id' => ?int,
  );

  const type TShape = shape(
    'major' => int,
    'package' => string,
    'annotation_with_quote' => string,
    'class_' => string,
    'annotation_with_trailing_comma' => string,
    'empty_annotations' => string,
    ?'my_enum' => ?MyEnum,
    'cpp_type_annotation' => vec<string>,
    ?'my_union' => ?MyUnion::TShape,
    'my_id' => int,
    ...
  );
  const int STRUCTURAL_ID = 2184994839913291136;
  /**
   * Original thrift field:-
   * 2: i64 major
   */
  public int $major;
  /**
   * Original thrift field:-
   * 1: string package
   */
  public string $package;
  /**
   * Original thrift field:-
   * 3: string annotation_with_quote
   */
  public string $annotation_with_quote;
  /**
   * Original thrift field:-
   * 4: string class_
   */
  public string $class_;
  /**
   * Original thrift field:-
   * 5: string annotation_with_trailing_comma
   */
  public string $annotation_with_trailing_comma;
  /**
   * Original thrift field:-
   * 6: string empty_annotations
   */
  public string $empty_annotations;
  /**
   * Original thrift field:-
   * 7: enum module.MyEnum my_enum
   */
  public ?MyEnum $my_enum;
  /**
   * Original thrift field:-
   * 8: list<string> cpp_type_annotation
   */
  public Vector<string> $cpp_type_annotation;
  /**
   * Original thrift field:-
   * 9: struct module.MyUnion my_union
   */
  public ?MyUnion $my_union;
  /**
   * Original thrift field:-
   * 10: module.MyId my_id
   */
  public int $my_id;

  public function __construct(?int $major = null, ?string $package = null, ?string $annotation_with_quote = null, ?string $class_ = null, ?string $annotation_with_trailing_comma = null, ?string $empty_annotations = null, ?MyEnum $my_enum = null, ?Vector<string> $cpp_type_annotation = null, ?MyUnion $my_union = null, ?int $my_id = null)[] {
    $this->major = $major ?? 0;
    $this->package = $package ?? '';
    $this->annotation_with_quote = $annotation_with_quote ?? '';
    $this->class_ = $class_ ?? '';
    $this->annotation_with_trailing_comma = $annotation_with_trailing_comma ?? '';
    $this->empty_annotations = $empty_annotations ?? '';
    $this->my_enum = $my_enum;
    $this->cpp_type_annotation = $cpp_type_annotation ?? Vector {};
    $this->my_union = $my_union;
    $this->my_id = $my_id ?? 0;
  }

  public static function withDefaultValues()[]: this {
    return new static();
  }

  public static function fromShape(self::TConstructorShape $shape)[]: this {
    return new static(
      Shapes::idx($shape, 'major'),
      Shapes::idx($shape, 'package'),
      Shapes::idx($shape, 'annotation_with_quote'),
      Shapes::idx($shape, 'class_'),
      Shapes::idx($shape, 'annotation_with_trailing_comma'),
      Shapes::idx($shape, 'empty_annotations'),
      Shapes::idx($shape, 'my_enum'),
      Shapes::idx($shape, 'cpp_type_annotation'),
      Shapes::idx($shape, 'my_union'),
      Shapes::idx($shape, 'my_id'),
    );
  }

  public function getName()[]: string {
    return 'MyStruct';
  }

  public function clearTerseFields()[write_props]: void {
  }

  public static function getStructMetadata()[]: \tmeta_ThriftStruct {
    return tmeta_ThriftStruct::fromShape(
      shape(
        "name" => "module.MyStruct",
        "fields" => vec[
          tmeta_ThriftField::fromShape(
            shape(
              "id" => 2,
              "type" => tmeta_ThriftType::fromShape(
                shape(
                  "t_primitive" => tmeta_ThriftPrimitiveType::THRIFT_I64_TYPE,
                )
              ),
              "name" => "major",
            )
          ),
          tmeta_ThriftField::fromShape(
            shape(
              "id" => 1,
              "type" => tmeta_ThriftType::fromShape(
                shape(
                  "t_primitive" => tmeta_ThriftPrimitiveType::THRIFT_STRING_TYPE,
                )
              ),
              "name" => "package",
            )
          ),
          tmeta_ThriftField::fromShape(
            shape(
              "id" => 3,
              "type" => tmeta_ThriftType::fromShape(
                shape(
                  "t_primitive" => tmeta_ThriftPrimitiveType::THRIFT_STRING_TYPE,
                )
              ),
              "name" => "annotation_with_quote",
            )
          ),
          tmeta_ThriftField::fromShape(
            shape(
              "id" => 4,
              "type" => tmeta_ThriftType::fromShape(
                shape(
                  "t_primitive" => tmeta_ThriftPrimitiveType::THRIFT_STRING_TYPE,
                )
              ),
              "name" => "class_",
            )
          ),
          tmeta_ThriftField::fromShape(
            shape(
              "id" => 5,
              "type" => tmeta_ThriftType::fromShape(
                shape(
                  "t_primitive" => tmeta_ThriftPrimitiveType::THRIFT_STRING_TYPE,
                )
              ),
              "name" => "annotation_with_trailing_comma",
            )
          ),
          tmeta_ThriftField::fromShape(
            shape(
              "id" => 6,
              "type" => tmeta_ThriftType::fromShape(
                shape(
                  "t_primitive" => tmeta_ThriftPrimitiveType::THRIFT_STRING_TYPE,
                )
              ),
              "name" => "empty_annotations",
            )
          ),
          tmeta_ThriftField::fromShape(
            shape(
              "id" => 7,
              "type" => tmeta_ThriftType::fromShape(
                shape(
                  "t_enum" => tmeta_ThriftEnumType::fromShape(
                    shape(
                      "name" => "module.MyEnum",
                    )
                  ),
                )
              ),
              "name" => "my_enum",
            )
          ),
          tmeta_ThriftField::fromShape(
            shape(
              "id" => 8,
              "type" => tmeta_ThriftType::fromShape(
                shape(
                  "t_list" => tmeta_ThriftListType::fromShape(
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
              "name" => "cpp_type_annotation",
            )
          ),
          tmeta_ThriftField::fromShape(
            shape(
              "id" => 9,
              "type" => tmeta_ThriftType::fromShape(
                shape(
                  "t_struct" => tmeta_ThriftStructType::fromShape(
                    shape(
                      "name" => "module.MyUnion",
                    )
                  ),
                )
              ),
              "name" => "my_union",
            )
          ),
          tmeta_ThriftField::fromShape(
            shape(
              "id" => 10,
              "type" => tmeta_ThriftType::fromShape(
                shape(
                  "t_typedef" => tmeta_ThriftTypedefType::fromShape(
                    shape(
                      "name" => "module.MyId",
                      "underlyingType" => tmeta_ThriftType::fromShape(
                        shape(
                          "t_primitive" => tmeta_ThriftPrimitiveType::THRIFT_I16_TYPE,
                        )
                      ),
                    )
                  ),
                )
              ),
              "name" => "my_id",
            )
          ),
        ],
        "is_union" => false,
      )
    );
  }

  public static function getAllStructuredAnnotations()[write_props]: \TStructAnnotations {
    return shape(
      'struct' => dict[
        '\thrift\annotation\cpp\Adapter' => \thrift\annotation\cpp\Adapter::fromShape(
          shape(
            "name" => "::StaticCast",
          )
        ),
      ],
      'fields' => dict[
        'my_union' => shape(
          'field' => dict[],
          'type' => dict[
            '\thrift\annotation\cpp\Adapter' => \thrift\annotation\cpp\Adapter::fromShape(
              shape(
                "name" => "::StaticCast",
              )
            ),
          ],
        ),
        'my_id' => shape(
          'field' => dict[],
          'type' => dict[
            '\thrift\annotation\cpp\StrongType' => \thrift\annotation\cpp\StrongType::fromShape(
              shape(
              )
            ),
          ],
        ),
      ],
    );
  }

  public static function __fromShape(self::TShape $shape)[]: this {
    return new static(
      $shape['major'],
      $shape['package'],
      $shape['annotation_with_quote'],
      $shape['class_'],
      $shape['annotation_with_trailing_comma'],
      $shape['empty_annotations'],
      Shapes::idx($shape, 'my_enum'),
      (new Vector($shape['cpp_type_annotation'])),
      Shapes::idx($shape, 'my_union') === null ? null : (MyUnion::__fromShape($shape['my_union'])),
      $shape['my_id'],
    );
  }

  public function __toShape()[]: self::TShape {
    return shape(
      'major' => $this->major,
      'package' => $this->package,
      'annotation_with_quote' => $this->annotation_with_quote,
      'class_' => $this->class_,
      'annotation_with_trailing_comma' => $this->annotation_with_trailing_comma,
      'empty_annotations' => $this->empty_annotations,
      'my_enum' => $this->my_enum,
      'cpp_type_annotation' => vec($this->cpp_type_annotation),
      'my_union' => $this->my_union?->__toShape(),
      'my_id' => $this->my_id,
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

    if (idx($parsed, 'major') !== null) {
      $this->major = HH\FIXME\UNSAFE_CAST<mixed, int>($parsed['major']);
    }
    if (idx($parsed, 'package') !== null) {
      $this->package = HH\FIXME\UNSAFE_CAST<mixed, string>($parsed['package']);
    }
    if (idx($parsed, 'annotation_with_quote') !== null) {
      $this->annotation_with_quote = HH\FIXME\UNSAFE_CAST<mixed, string>($parsed['annotation_with_quote']);
    }
    if (idx($parsed, 'class_') !== null) {
      $this->class_ = HH\FIXME\UNSAFE_CAST<mixed, string>($parsed['class_']);
    }
    if (idx($parsed, 'annotation_with_trailing_comma') !== null) {
      $this->annotation_with_trailing_comma = HH\FIXME\UNSAFE_CAST<mixed, string>($parsed['annotation_with_trailing_comma']);
    }
    if (idx($parsed, 'empty_annotations') !== null) {
      $this->empty_annotations = HH\FIXME\UNSAFE_CAST<mixed, string>($parsed['empty_annotations']);
    }
    if (idx($parsed, 'my_enum') !== null) {
      $this->my_enum = MyEnum::coerce(HH\FIXME\UNSAFE_CAST<mixed, MyEnum>($parsed['my_enum']));
    }
    if (idx($parsed, 'cpp_type_annotation') !== null) {
      $_json3 = HH\FIXME\UNSAFE_CAST<mixed, Vector<string>>($parsed['cpp_type_annotation']);
      $_container4 = Vector {};
      foreach($_json3 as $_key1 => $_value2) {
        $_elem5 = '';
        $_elem5 = $_value2;
        $_container4 []= $_elem5;
      }
      $this->cpp_type_annotation = $_container4;
    }
    if (idx($parsed, 'my_union') !== null) {
      $_tmp6 = \json_encode(HH\FIXME\UNSAFE_CAST<mixed, MyUnion>($parsed['my_union']));
      $_tmp7 = MyUnion::withDefaultValues();
      $_tmp7->readFromJson($_tmp6);
      $this->my_union = $_tmp7;
    }
    if (idx($parsed, 'my_id') !== null) {
      $_tmp8 = (int)HH\FIXME\UNSAFE_CAST<mixed, int>($parsed['my_id']);
      if ($_tmp8 > 0x7fff) {
        throw new \TProtocolException("number exceeds limit in field");
      } else {
        $this->my_id = (int)$_tmp8;
      }
    }
  }

}

/**
 * Original thrift struct:-
 * SecretStruct
 */
class SecretStruct implements \IThriftSyncStruct, \IThriftStructMetadata, \IThriftShapishSyncStruct {
  use \ThriftSerializationTrait;

  const \ThriftStructTypes::TSpec SPEC = dict[
    1 => shape(
      'var' => 'id',
      'type' => \TType::I64,
    ),
    2 => shape(
      'var' => 'password',
      'type' => \TType::STRING,
    ),
  ];
  const dict<string, int> FIELDMAP = dict[
    'id' => 1,
    'password' => 2,
  ];

  const type TConstructorShape = shape(
    ?'id' => ?int,
    ?'password' => ?string,
  );

  const type TShape = shape(
    'id' => int,
    'password' => string,
    ...
  );
  const int STRUCTURAL_ID = 7563936987719176841;
  /**
   * Original thrift field:-
   * 1: i64 id
   */
  public int $id;
  /**
   * Original thrift field:-
   * 2: string password
   */
  public string $password;

  public function __construct(?int $id = null, ?string $password = null)[] {
    $this->id = $id ?? 0;
    $this->password = $password ?? '';
  }

  public static function withDefaultValues()[]: this {
    return new static();
  }

  public static function fromShape(self::TConstructorShape $shape)[]: this {
    return new static(
      Shapes::idx($shape, 'id'),
      Shapes::idx($shape, 'password'),
    );
  }

  public function getName()[]: string {
    return 'SecretStruct';
  }

  public function clearTerseFields()[write_props]: void {
  }

  public static function getStructMetadata()[]: \tmeta_ThriftStruct {
    return tmeta_ThriftStruct::fromShape(
      shape(
        "name" => "module.SecretStruct",
        "fields" => vec[
          tmeta_ThriftField::fromShape(
            shape(
              "id" => 1,
              "type" => tmeta_ThriftType::fromShape(
                shape(
                  "t_primitive" => tmeta_ThriftPrimitiveType::THRIFT_I64_TYPE,
                )
              ),
              "name" => "id",
            )
          ),
          tmeta_ThriftField::fromShape(
            shape(
              "id" => 2,
              "type" => tmeta_ThriftType::fromShape(
                shape(
                  "t_primitive" => tmeta_ThriftPrimitiveType::THRIFT_STRING_TYPE,
                )
              ),
              "name" => "password",
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

  public static function __fromShape(self::TShape $shape)[]: this {
    return new static(
      $shape['id'],
      $shape['password'],
    );
  }

  public function __toShape()[]: self::TShape {
    return shape(
      'id' => $this->id,
      'password' => $this->password,
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

    if (idx($parsed, 'id') !== null) {
      $this->id = HH\FIXME\UNSAFE_CAST<mixed, int>($parsed['id']);
    }
    if (idx($parsed, 'password') !== null) {
      $this->password = HH\FIXME\UNSAFE_CAST<mixed, string>($parsed['password']);
    }
  }

}

