/**
 * Autogenerated by Thrift
 *
 * DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING
 *  @generated
 */
package com.facebook.thrift.annotation.java_deprecated;

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
public class Annotation implements TBase, java.io.Serializable, Cloneable, Comparable<Annotation> {
  private static final TStruct STRUCT_DESC = new TStruct("Annotation");
  private static final TField JAVA_ANNOTATION_FIELD_DESC = new TField("java_annotation", TType.STRING, (short)1);

  public String java_annotation;
  public static final int JAVA_ANNOTATION = 1;

  // isset id assignments

  public static final Map<Integer, FieldMetaData> metaDataMap;

  static {
    Map<Integer, FieldMetaData> tmpMetaDataMap = new HashMap<Integer, FieldMetaData>();
    tmpMetaDataMap.put(JAVA_ANNOTATION, new FieldMetaData("java_annotation", TFieldRequirementType.DEFAULT, 
        new FieldValueMetaData(TType.STRING)));
    metaDataMap = Collections.unmodifiableMap(tmpMetaDataMap);
  }

  static {
    FieldMetaData.addStructMetaDataMap(Annotation.class, metaDataMap);
  }

  public Annotation() {
  }

  public Annotation(
      String java_annotation) {
    this();
    this.java_annotation = java_annotation;
  }

  public static class Builder {
    private String java_annotation;

    public Builder() {
    }

    public Builder setJava_annotation(final String java_annotation) {
      this.java_annotation = java_annotation;
      return this;
    }

    public Annotation build() {
      Annotation result = new Annotation();
      result.setJava_annotation(this.java_annotation);
      return result;
    }
  }

  public static Builder builder() {
    return new Builder();
  }

  /**
   * Performs a deep copy on <i>other</i>.
   */
  public Annotation(Annotation other) {
    if (other.isSetJava_annotation()) {
      this.java_annotation = TBaseHelper.deepCopy(other.java_annotation);
    }
  }

  public Annotation deepCopy() {
    return new Annotation(this);
  }

  public String getJava_annotation() {
    return this.java_annotation;
  }

  public Annotation setJava_annotation(String java_annotation) {
    this.java_annotation = java_annotation;
    return this;
  }

  public void unsetJava_annotation() {
    this.java_annotation = null;
  }

  // Returns true if field java_annotation is set (has been assigned a value) and false otherwise
  public boolean isSetJava_annotation() {
    return this.java_annotation != null;
  }

  public void setJava_annotationIsSet(boolean __value) {
    if (!__value) {
      this.java_annotation = null;
    }
  }

  public void setFieldValue(int fieldID, Object __value) {
    switch (fieldID) {
    case JAVA_ANNOTATION:
      if (__value == null) {
        unsetJava_annotation();
      } else {
        setJava_annotation((String)__value);
      }
      break;

    default:
      throw new IllegalArgumentException("Field " + fieldID + " doesn't exist!");
    }
  }

  public Object getFieldValue(int fieldID) {
    switch (fieldID) {
    case JAVA_ANNOTATION:
      return getJava_annotation();

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
    if (!(_that instanceof Annotation))
      return false;
    Annotation that = (Annotation)_that;

    if (!TBaseHelper.equalsNobinary(this.isSetJava_annotation(), that.isSetJava_annotation(), this.java_annotation, that.java_annotation)) { return false; }

    return true;
  }

  @Override
  public int hashCode() {
    return Arrays.deepHashCode(new Object[] {java_annotation});
  }

  @Override
  public int compareTo(Annotation other) {
    if (other == null) {
      // See java.lang.Comparable docs
      throw new NullPointerException();
    }

    if (other == this) {
      return 0;
    }
    int lastComparison = 0;

    lastComparison = Boolean.valueOf(isSetJava_annotation()).compareTo(other.isSetJava_annotation());
    if (lastComparison != 0) {
      return lastComparison;
    }
    lastComparison = TBaseHelper.compareTo(java_annotation, other.java_annotation);
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
        case JAVA_ANNOTATION:
          if (__field.type == TType.STRING) {
            this.java_annotation = iprot.readString();
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
    if (this.java_annotation != null) {
      oprot.writeFieldBegin(JAVA_ANNOTATION_FIELD_DESC);
      oprot.writeString(this.java_annotation);
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
    StringBuilder sb = new StringBuilder("Annotation");
    sb.append(space);
    sb.append("(");
    sb.append(newLine);
    boolean first = true;

    sb.append(indentStr);
    sb.append("java_annotation");
    sb.append(space);
    sb.append(":").append(space);
    if (this.getJava_annotation() == null) {
      sb.append("null");
    } else {
      sb.append(TBaseHelper.toString(this.getJava_annotation(), indent + 1, prettyPrint));
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
