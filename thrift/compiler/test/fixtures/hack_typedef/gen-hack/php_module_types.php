<?hh
/**
 * Autogenerated by Thrift
 *
 * DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING
 *  @generated
 */

type include_typedef_MyI32 = int;
/**
 * Original thrift struct:-
 * Foo
 */
class include_typedef_Foo implements \IThriftSyncStruct, \IThriftStructMetadata {
  use \ThriftSerializationTrait;

  const \ThriftStructTypes::TSpec SPEC = dict[
    1 => shape(
      'var' => 'i_field',
      'type' => \TType::I32,
    ),
    2 => shape(
      'var' => 'i_field2',
      'type' => \TType::I32,
    ),
  ];
  const dict<string, int> FIELDMAP = dict[
    'i_field' => 1,
    'i_field2' => 2,
  ];

  const type TConstructorShape = shape(
    ?'i_field' => ?include_typedef_MyI32,
    ?'i_field2' => ?int,
  );

  const int STRUCTURAL_ID = 4034280781637513039;
  /**
   * Original thrift field:-
   * 1: php_module.MyI32 i_field
   */
  public include_typedef_MyI32 $i_field;
  /**
   * Original thrift field:-
   * 2: i32 i_field2
   */
  public int $i_field2;

  public function __construct(?include_typedef_MyI32 $i_field = null, ?int $i_field2 = null)[] {
    $this->i_field = $i_field ?? 0;
    $this->i_field2 = $i_field2 ?? 0;
  }

  public static function withDefaultValues()[]: this {
    return new static();
  }

  public static function fromShape(self::TConstructorShape $shape)[]: this {
    return new static(
      Shapes::idx($shape, 'i_field'),
      Shapes::idx($shape, 'i_field2'),
    );
  }

  public function getName()[]: string {
    return 'Foo';
  }

  public function clearTerseFields()[write_props]: void {
  }

  public static function getStructMetadata()[]: \tmeta_ThriftStruct {
    return tmeta_ThriftStruct::fromShape(
      shape(
        "name" => "php_module.Foo",
        "fields" => vec[
          tmeta_ThriftField::fromShape(
            shape(
              "id" => 1,
              "type" => tmeta_ThriftType::fromShape(
                shape(
                  "t_typedef" => tmeta_ThriftTypedefType::fromShape(
                    shape(
                      "name" => "php_module.MyI32",
                      "underlyingType" => tmeta_ThriftType::fromShape(
                        shape(
                          "t_primitive" => tmeta_ThriftPrimitiveType::THRIFT_I32_TYPE,
                        )
                      ),
                    )
                  ),
                )
              ),
              "name" => "i_field",
            )
          ),
          tmeta_ThriftField::fromShape(
            shape(
              "id" => 2,
              "type" => tmeta_ThriftType::fromShape(
                shape(
                  "t_primitive" => tmeta_ThriftPrimitiveType::THRIFT_I32_TYPE,
                )
              ),
              "name" => "i_field2",
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

