/**
 * Autogenerated by Thrift
 *
 * DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING
 *  @generated
 */
package test.fixtures.basicannotations;

import java.util.List;
import java.util.ArrayList;
import java.util.Map;
import java.util.HashMap;
import java.util.Set;
import java.util.HashSet;
import java.util.Collections;
import java.util.BitSet;
import java.util.Arrays;
import com.facebook.thrift.*;
import com.facebook.thrift.annotations.*;
import com.facebook.thrift.async.*;
import com.facebook.thrift.meta_data.*;
import com.facebook.thrift.server.*;
import com.facebook.thrift.transport.*;
import com.facebook.thrift.protocol.*;

@SuppressWarnings({ "unused", "serial" })
public class MyStructAnnotation implements TBase, java.io.Serializable, Cloneable, Comparable<MyStructAnnotation> {
  private static final TStruct STRUCT_DESC = new TStruct("MyStructAnnotation");
  private static final TField COUNT_FIELD_DESC = new TField("count", TType.I64, (short)1);
  private static final TField NAME_FIELD_DESC = new TField("name", TType.STRING, (short)2);
  private static final TField EXTRA_FIELD_DESC = new TField("extra", TType.STRING, (short)3);
  private static final TField NEST_FIELD_DESC = new TField("nest", TType.STRUCT, (short)4);

  public long count;
  public String name;
  public String extra;
  public MyStructNestedAnnotation nest;
  public static final int COUNT = 1;
  public static final int NAME = 2;
  public static final int EXTRA = 3;
  public static final int NEST = 4;

  // isset id assignments
  private static final int __COUNT_ISSET_ID = 0;
  private BitSet __isset_bit_vector = new BitSet(1);

  public static final Map<Integer, FieldMetaData> metaDataMap;

  static {
    Map<Integer, FieldMetaData> tmpMetaDataMap = new HashMap<Integer, FieldMetaData>();
    tmpMetaDataMap.put(COUNT, new FieldMetaData("count", TFieldRequirementType.DEFAULT, 
        new FieldValueMetaData(TType.I64)));
    tmpMetaDataMap.put(NAME, new FieldMetaData("name", TFieldRequirementType.DEFAULT, 
        new FieldValueMetaData(TType.STRING)));
    tmpMetaDataMap.put(EXTRA, new FieldMetaData("extra", TFieldRequirementType.OPTIONAL, 
        new FieldValueMetaData(TType.STRING)));
    tmpMetaDataMap.put(NEST, new FieldMetaData("nest", TFieldRequirementType.DEFAULT, 
        new StructMetaData(TType.STRUCT, MyStructNestedAnnotation.class)));
    metaDataMap = Collections.unmodifiableMap(tmpMetaDataMap);
  }

  static {
    FieldMetaData.addStructMetaDataMap(MyStructAnnotation.class, metaDataMap);
  }

  public MyStructAnnotation() {
  }

  public MyStructAnnotation(
    long count,
    String name,
    MyStructNestedAnnotation nest)
  {
    this();
    this.count = count;
    setCountIsSet(true);
    this.name = name;
    this.nest = nest;
  }

  public MyStructAnnotation(
    long count,
    String name,
    String extra,
    MyStructNestedAnnotation nest)
  {
    this();
    this.count = count;
    setCountIsSet(true);
    this.name = name;
    this.extra = extra;
    this.nest = nest;
  }

  public static class Builder {
    private long count;
    private String name;
    private String extra;
    private MyStructNestedAnnotation nest;

    BitSet __optional_isset = new BitSet(1);

    public Builder() {
    }

    public Builder setCount(final long count) {
      this.count = count;
      __optional_isset.set(__COUNT_ISSET_ID, true);
      return this;
    }

    public Builder setName(final String name) {
      this.name = name;
      return this;
    }

    public Builder setExtra(final String extra) {
      this.extra = extra;
      return this;
    }

    public Builder setNest(final MyStructNestedAnnotation nest) {
      this.nest = nest;
      return this;
    }

    public MyStructAnnotation build() {
      MyStructAnnotation result = new MyStructAnnotation();
      if (__optional_isset.get(__COUNT_ISSET_ID)) {
        result.setCount(this.count);
      }
      result.setName(this.name);
      result.setExtra(this.extra);
      result.setNest(this.nest);
      return result;
    }
  }

  public static Builder builder() {
    return new Builder();
  }

  /**
   * Performs a deep copy on <i>other</i>.
   */
  public MyStructAnnotation(MyStructAnnotation other) {
    __isset_bit_vector.clear();
    __isset_bit_vector.or(other.__isset_bit_vector);
    this.count = TBaseHelper.deepCopy(other.count);
    if (other.isSetName()) {
      this.name = TBaseHelper.deepCopy(other.name);
    }
    if (other.isSetExtra()) {
      this.extra = TBaseHelper.deepCopy(other.extra);
    }
    if (other.isSetNest()) {
      this.nest = TBaseHelper.deepCopy(other.nest);
    }
  }

  public MyStructAnnotation deepCopy() {
    return new MyStructAnnotation(this);
  }

  public long getCount() {
    return this.count;
  }

  public MyStructAnnotation setCount(long count) {
    this.count = count;
    setCountIsSet(true);
    return this;
  }

  public void unsetCount() {
    __isset_bit_vector.clear(__COUNT_ISSET_ID);
  }

  // Returns true if field count is set (has been assigned a value) and false otherwise
  public boolean isSetCount() {
    return __isset_bit_vector.get(__COUNT_ISSET_ID);
  }

  public void setCountIsSet(boolean __value) {
    __isset_bit_vector.set(__COUNT_ISSET_ID, __value);
  }

  public String getName() {
    return this.name;
  }

  public MyStructAnnotation setName(String name) {
    this.name = name;
    return this;
  }

  public void unsetName() {
    this.name = null;
  }

  // Returns true if field name is set (has been assigned a value) and false otherwise
  public boolean isSetName() {
    return this.name != null;
  }

  public void setNameIsSet(boolean __value) {
    if (!__value) {
      this.name = null;
    }
  }

  public String getExtra() {
    return this.extra;
  }

  public MyStructAnnotation setExtra(String extra) {
    this.extra = extra;
    return this;
  }

  public void unsetExtra() {
    this.extra = null;
  }

  // Returns true if field extra is set (has been assigned a value) and false otherwise
  public boolean isSetExtra() {
    return this.extra != null;
  }

  public void setExtraIsSet(boolean __value) {
    if (!__value) {
      this.extra = null;
    }
  }

  public MyStructNestedAnnotation getNest() {
    return this.nest;
  }

  public MyStructAnnotation setNest(MyStructNestedAnnotation nest) {
    this.nest = nest;
    return this;
  }

  public void unsetNest() {
    this.nest = null;
  }

  // Returns true if field nest is set (has been assigned a value) and false otherwise
  public boolean isSetNest() {
    return this.nest != null;
  }

  public void setNestIsSet(boolean __value) {
    if (!__value) {
      this.nest = null;
    }
  }

  public void setFieldValue(int fieldID, Object __value) {
    switch (fieldID) {
    case COUNT:
      if (__value == null) {
        unsetCount();
      } else {
        setCount((Long)__value);
      }
      break;

    case NAME:
      if (__value == null) {
        unsetName();
      } else {
        setName((String)__value);
      }
      break;

    case EXTRA:
      if (__value == null) {
        unsetExtra();
      } else {
        setExtra((String)__value);
      }
      break;

    case NEST:
      if (__value == null) {
        unsetNest();
      } else {
        setNest((MyStructNestedAnnotation)__value);
      }
      break;

    default:
      throw new IllegalArgumentException("Field " + fieldID + " doesn't exist!");
    }
  }

  public Object getFieldValue(int fieldID) {
    switch (fieldID) {
    case COUNT:
      return new Long(getCount());

    case NAME:
      return getName();

    case EXTRA:
      return getExtra();

    case NEST:
      return getNest();

    default:
      throw new IllegalArgumentException("Field " + fieldID + " doesn't exist!");
    }
  }

  @Override
  public boolean equals(Object _that) {
    if (_that == null)
      return false;
    if (this == _that)
      return true;
    if (!(_that instanceof MyStructAnnotation))
      return false;
    MyStructAnnotation that = (MyStructAnnotation)_that;

    if (!TBaseHelper.equalsNobinary(this.count, that.count)) { return false; }

    if (!TBaseHelper.equalsNobinary(this.isSetName(), that.isSetName(), this.name, that.name)) { return false; }

    if (!TBaseHelper.equalsNobinary(this.isSetExtra(), that.isSetExtra(), this.extra, that.extra)) { return false; }

    if (!TBaseHelper.equalsNobinary(this.isSetNest(), that.isSetNest(), this.nest, that.nest)) { return false; }

    return true;
  }

  @Override
  public int hashCode() {
    return Arrays.deepHashCode(new Object[] {count, name, extra, nest});
  }

  @Override
  public int compareTo(MyStructAnnotation other) {
    if (other == null) {
      // See java.lang.Comparable docs
      throw new NullPointerException();
    }

    if (other == this) {
      return 0;
    }
    int lastComparison = 0;

    lastComparison = Boolean.valueOf(isSetCount()).compareTo(other.isSetCount());
    if (lastComparison != 0) {
      return lastComparison;
    }
    lastComparison = TBaseHelper.compareTo(count, other.count);
    if (lastComparison != 0) { 
      return lastComparison;
    }
    lastComparison = Boolean.valueOf(isSetName()).compareTo(other.isSetName());
    if (lastComparison != 0) {
      return lastComparison;
    }
    lastComparison = TBaseHelper.compareTo(name, other.name);
    if (lastComparison != 0) { 
      return lastComparison;
    }
    lastComparison = Boolean.valueOf(isSetExtra()).compareTo(other.isSetExtra());
    if (lastComparison != 0) {
      return lastComparison;
    }
    lastComparison = TBaseHelper.compareTo(extra, other.extra);
    if (lastComparison != 0) { 
      return lastComparison;
    }
    lastComparison = Boolean.valueOf(isSetNest()).compareTo(other.isSetNest());
    if (lastComparison != 0) {
      return lastComparison;
    }
    lastComparison = TBaseHelper.compareTo(nest, other.nest);
    if (lastComparison != 0) { 
      return lastComparison;
    }
    return 0;
  }

  public void read(TProtocol iprot) throws TException {
    TField __field;
    iprot.readStructBegin(metaDataMap);
    while (true)
    {
      __field = iprot.readFieldBegin();
      if (__field.type == TType.STOP) { 
        break;
      }
      switch (__field.id)
      {
        case COUNT:
          if (__field.type == TType.I64) {
            this.count = iprot.readI64();
            setCountIsSet(true);
          } else { 
            TProtocolUtil.skip(iprot, __field.type);
          }
          break;
        case NAME:
          if (__field.type == TType.STRING) {
            this.name = iprot.readString();
          } else { 
            TProtocolUtil.skip(iprot, __field.type);
          }
          break;
        case EXTRA:
          if (__field.type == TType.STRING) {
            this.extra = iprot.readString();
          } else { 
            TProtocolUtil.skip(iprot, __field.type);
          }
          break;
        case NEST:
          if (__field.type == TType.STRUCT) {
            this.nest = new MyStructNestedAnnotation();
            this.nest.read(iprot);
          } else { 
            TProtocolUtil.skip(iprot, __field.type);
          }
          break;
        default:
          TProtocolUtil.skip(iprot, __field.type);
          break;
      }
      iprot.readFieldEnd();
    }
    iprot.readStructEnd();


    // check for required fields of primitive type, which can't be checked in the validate method
    validate();
  }

  public void write(TProtocol oprot) throws TException {
    validate();

    oprot.writeStructBegin(STRUCT_DESC);
    oprot.writeFieldBegin(COUNT_FIELD_DESC);
    oprot.writeI64(this.count);
    oprot.writeFieldEnd();
    if (this.name != null) {
      oprot.writeFieldBegin(NAME_FIELD_DESC);
      oprot.writeString(this.name);
      oprot.writeFieldEnd();
    }
    if (this.extra != null) {
      if (isSetExtra()) {
        oprot.writeFieldBegin(EXTRA_FIELD_DESC);
        oprot.writeString(this.extra);
        oprot.writeFieldEnd();
      }
    }
    if (this.nest != null) {
      oprot.writeFieldBegin(NEST_FIELD_DESC);
      this.nest.write(oprot);
      oprot.writeFieldEnd();
    }
    oprot.writeFieldStop();
    oprot.writeStructEnd();
  }

  @Override
  public String toString() {
    return toString(1, true);
  }

  @Override
  public String toString(int indent, boolean prettyPrint) {
    String indentStr = prettyPrint ? TBaseHelper.getIndentedString(indent) : "";
    String newLine = prettyPrint ? "\n" : "";
    String space = prettyPrint ? " " : "";
    StringBuilder sb = new StringBuilder("MyStructAnnotation");
    sb.append(space);
    sb.append("(");
    sb.append(newLine);
    boolean first = true;

    sb.append(indentStr);
    sb.append("count");
    sb.append(space);
    sb.append(":").append(space);
    sb.append(TBaseHelper.toString(this.getCount(), indent + 1, prettyPrint));
    first = false;
    if (!first) sb.append("," + newLine);
    sb.append(indentStr);
    sb.append("name");
    sb.append(space);
    sb.append(":").append(space);
    if (this.getName() == null) {
      sb.append("null");
    } else {
      sb.append(TBaseHelper.toString(this.getName(), indent + 1, prettyPrint));
    }
    first = false;
    if (isSetExtra())
    {
      if (!first) sb.append("," + newLine);
      sb.append(indentStr);
      sb.append("extra");
      sb.append(space);
      sb.append(":").append(space);
      if (this.getExtra() == null) {
        sb.append("null");
      } else {
        sb.append(TBaseHelper.toString(this.getExtra(), indent + 1, prettyPrint));
      }
      first = false;
    }
    if (!first) sb.append("," + newLine);
    sb.append(indentStr);
    sb.append("nest");
    sb.append(space);
    sb.append(":").append(space);
    if (this.getNest() == null) {
      sb.append("null");
    } else {
      sb.append(TBaseHelper.toString(this.getNest(), indent + 1, prettyPrint));
    }
    first = false;
    sb.append(newLine + TBaseHelper.reduceIndent(indentStr));
    sb.append(")");
    return sb.toString();
  }

  public void validate() throws TException {
    // check for required fields
  }

}

