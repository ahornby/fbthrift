/**
 * Autogenerated by Thrift
 *
 * DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING
 *  @generated
 */

package test.fixtures.patch;

import com.facebook.swift.codec.*;
import com.facebook.swift.codec.ThriftField.Requiredness;
import com.facebook.swift.codec.ThriftField.Recursiveness;
import com.google.common.collect.*;
import java.util.*;
import javax.annotation.Nullable;
import org.apache.thrift.*;
import org.apache.thrift.transport.*;
import org.apache.thrift.protocol.*;
import static com.google.common.base.MoreObjects.toStringHelper;
import static com.google.common.base.MoreObjects.ToStringHelper;

@SwiftGenerated
@com.facebook.swift.codec.ThriftStruct(value="RefFieldsField5Patch", builder=RefFieldsField5Patch.Builder.class)
public final class RefFieldsField5Patch implements com.facebook.thrift.payload.ThriftSerializable {
    @ThriftConstructor
    public RefFieldsField5Patch(
        @com.facebook.swift.codec.ThriftField(value=1, name="assign", requiredness=Requiredness.OPTIONAL) final List<Integer> assign,
        @com.facebook.swift.codec.ThriftField(value=2, name="clear", requiredness=Requiredness.TERSE) final boolean clear,
        @com.facebook.swift.codec.ThriftField(value=8, name="prepend", requiredness=Requiredness.TERSE) final List<Integer> prepend,
        @com.facebook.swift.codec.ThriftField(value=9, name="append", requiredness=Requiredness.TERSE) final List<Integer> append
    ) {
        this.assign = assign;
        this.clear = clear;
        this.prepend = prepend;
        this.append = append;
    }
    
    @ThriftConstructor
    protected RefFieldsField5Patch() {
      this.assign = null;
      this.clear = false;
      this.prepend = com.facebook.thrift.util.IntrinsicDefaults.defaultList();
      this.append = com.facebook.thrift.util.IntrinsicDefaults.defaultList();
    }
    
    public static class Builder {
        private List<Integer> assign = null;
        private boolean clear = false;
        private List<Integer> prepend = com.facebook.thrift.util.IntrinsicDefaults.defaultList();
        private List<Integer> append = com.facebook.thrift.util.IntrinsicDefaults.defaultList();
    
        @com.facebook.swift.codec.ThriftField(value=1, name="assign", requiredness=Requiredness.OPTIONAL)
        public Builder setAssign(List<Integer> assign) {
            this.assign = assign;
            return this;
        }
    
        public List<Integer> getAssign() { return assign; }
    
            @com.facebook.swift.codec.ThriftField(value=2, name="clear", requiredness=Requiredness.TERSE)
        public Builder setClear(boolean clear) {
            this.clear = clear;
            return this;
        }
    
        public boolean isClear() { return clear; }
    
            @com.facebook.swift.codec.ThriftField(value=8, name="prepend", requiredness=Requiredness.TERSE)
        public Builder setPrepend(List<Integer> prepend) {
            this.prepend = prepend;
            return this;
        }
    
        public List<Integer> getPrepend() { return prepend; }
    
            @com.facebook.swift.codec.ThriftField(value=9, name="append", requiredness=Requiredness.TERSE)
        public Builder setAppend(List<Integer> append) {
            this.append = append;
            return this;
        }
    
        public List<Integer> getAppend() { return append; }
    
        public Builder() { }
        public Builder(RefFieldsField5Patch other) {
            this.assign = other.assign;
            this.clear = other.clear;
            this.prepend = other.prepend;
            this.append = other.append;
        }
    
        @ThriftConstructor
        public RefFieldsField5Patch build() {
            RefFieldsField5Patch result = new RefFieldsField5Patch (
                this.assign,
                this.clear,
                this.prepend,
                this.append
            );
            return result;
        }
    }
        
    public static final Map<String, Integer> NAMES_TO_IDS = new HashMap();
    public static final Map<String, Integer> THRIFT_NAMES_TO_IDS = new HashMap();
    public static final Map<Integer, TField> FIELD_METADATA = new HashMap<>();
    private static final TStruct STRUCT_DESC = new TStruct("RefFieldsField5Patch");
    private final List<Integer> assign;
    public static final int _ASSIGN = 1;
    private static final TField ASSIGN_FIELD_DESC = new TField("assign", TType.LIST, (short)1);
        private final boolean clear;
    public static final int _CLEAR = 2;
    private static final TField CLEAR_FIELD_DESC = new TField("clear", TType.BOOL, (short)2);
        private final List<Integer> prepend;
    public static final int _PREPEND = 8;
    private static final TField PREPEND_FIELD_DESC = new TField("prepend", TType.LIST, (short)8);
        private final List<Integer> append;
    public static final int _APPEND = 9;
    private static final TField APPEND_FIELD_DESC = new TField("append", TType.LIST, (short)9);
    static {
      NAMES_TO_IDS.put("assign", 1);
      THRIFT_NAMES_TO_IDS.put("assign", 1);
      FIELD_METADATA.put(1, ASSIGN_FIELD_DESC);
      NAMES_TO_IDS.put("clear", 2);
      THRIFT_NAMES_TO_IDS.put("clear", 2);
      FIELD_METADATA.put(2, CLEAR_FIELD_DESC);
      NAMES_TO_IDS.put("prepend", 8);
      THRIFT_NAMES_TO_IDS.put("prepend", 8);
      FIELD_METADATA.put(8, PREPEND_FIELD_DESC);
      NAMES_TO_IDS.put("append", 9);
      THRIFT_NAMES_TO_IDS.put("append", 9);
      FIELD_METADATA.put(9, APPEND_FIELD_DESC);
      com.facebook.thrift.type.TypeRegistry.add(new com.facebook.thrift.type.Type(
        new com.facebook.thrift.type.UniversalName("test.dev/fixtures/patch/RefFieldsField5Patch"),
        RefFieldsField5Patch.class, RefFieldsField5Patch::read0));
    }
    
    @Nullable
    @com.facebook.swift.codec.ThriftField(value=1, name="assign", requiredness=Requiredness.OPTIONAL)
    public List<Integer> getAssign() { return assign; }
    
    
    
    @com.facebook.swift.codec.ThriftField(value=2, name="clear", requiredness=Requiredness.TERSE)
    public boolean isClear() { return clear; }
    
    
    @Nullable
    @com.facebook.swift.codec.ThriftField(value=8, name="prepend", requiredness=Requiredness.TERSE)
    public List<Integer> getPrepend() { return prepend; }
    
    
    @Nullable
    @com.facebook.swift.codec.ThriftField(value=9, name="append", requiredness=Requiredness.TERSE)
    public List<Integer> getAppend() { return append; }
    
    @java.lang.Override
    public String toString() {
        ToStringHelper helper = toStringHelper(this);
        helper.add("assign", assign);
        helper.add("clear", clear);
        helper.add("prepend", prepend);
        helper.add("append", append);
        return helper.toString();
    }
    
    @java.lang.Override
    public boolean equals(java.lang.Object o) {
        if (this == o) {
            return true;
        }
        if (o == null || getClass() != o.getClass()) {
            return false;
        }
    
        RefFieldsField5Patch other = (RefFieldsField5Patch)o;
    
        return
            Objects.equals(assign, other.assign) &&
            Objects.equals(clear, other.clear) &&
            Objects.equals(prepend, other.prepend) &&
            Objects.equals(append, other.append) &&
            true;
    }
    
    @java.lang.Override
    public int hashCode() {
        return Arrays.deepHashCode(new java.lang.Object[] {
            assign,
            clear,
            prepend,
            append
        });
    }
    
    
    public static com.facebook.thrift.payload.Reader<RefFieldsField5Patch> asReader() {
      return RefFieldsField5Patch::read0;
    }
    
    public static RefFieldsField5Patch read0(TProtocol oprot) throws TException {
      TField __field;
      oprot.readStructBegin(RefFieldsField5Patch.NAMES_TO_IDS, RefFieldsField5Patch.THRIFT_NAMES_TO_IDS, RefFieldsField5Patch.FIELD_METADATA);
      RefFieldsField5Patch.Builder builder = new RefFieldsField5Patch.Builder();
      while (true) {
        __field = oprot.readFieldBegin();
        if (__field.type == TType.STOP) { break; }
        switch (__field.id) {
        case _ASSIGN:
          if (__field.type == TType.LIST) {
            List<Integer> assign;
            {
            TList _list = oprot.readListBegin();
            assign = new ArrayList<Integer>(Math.max(0, _list.size));
            for (int _i = 0; (_list.size < 0) ? oprot.peekList() : (_i < _list.size); _i++) {
                
                int _value1 = oprot.readI32();
                assign.add(_value1);
            }
            oprot.readListEnd();
            }
            builder.setAssign(assign);
          } else {
            TProtocolUtil.skip(oprot, __field.type);
          }
          break;
        case _CLEAR:
          if (__field.type == TType.BOOL) {
            boolean clear = oprot.readBool();
            builder.setClear(clear);
          } else {
            TProtocolUtil.skip(oprot, __field.type);
          }
          break;
        case _PREPEND:
          if (__field.type == TType.LIST) {
            List<Integer> prepend;
            {
            TList _list = oprot.readListBegin();
            prepend = new ArrayList<Integer>(Math.max(0, _list.size));
            for (int _i = 0; (_list.size < 0) ? oprot.peekList() : (_i < _list.size); _i++) {
                
                int _value1 = oprot.readI32();
                prepend.add(_value1);
            }
            oprot.readListEnd();
            }
            builder.setPrepend(prepend);
          } else {
            TProtocolUtil.skip(oprot, __field.type);
          }
          break;
        case _APPEND:
          if (__field.type == TType.LIST) {
            List<Integer> append;
            {
            TList _list = oprot.readListBegin();
            append = new ArrayList<Integer>(Math.max(0, _list.size));
            for (int _i = 0; (_list.size < 0) ? oprot.peekList() : (_i < _list.size); _i++) {
                
                int _value1 = oprot.readI32();
                append.add(_value1);
            }
            oprot.readListEnd();
            }
            builder.setAppend(append);
          } else {
            TProtocolUtil.skip(oprot, __field.type);
          }
          break;
        default:
          TProtocolUtil.skip(oprot, __field.type);
          break;
        }
        oprot.readFieldEnd();
      }
      oprot.readStructEnd();
      return builder.build();
    }
    
    public void write0(TProtocol oprot) throws TException {
      oprot.writeStructBegin(STRUCT_DESC);
      int structStart = 0;
      int pos = 0;
      com.facebook.thrift.protocol.ByteBufTProtocol p = (com.facebook.thrift.protocol.ByteBufTProtocol) oprot;
      if (assign != null) {
        oprot.writeFieldBegin(ASSIGN_FIELD_DESC);
        List<Integer> _iter0 = assign;
        oprot.writeListBegin(new TList(TType.I32, _iter0.size()));
        for (int _iter1 : _iter0) {
          oprot.writeI32(_iter1);
        }
        oprot.writeListEnd();
        oprot.writeFieldEnd();
      }
      if (!com.facebook.thrift.util.IntrinsicDefaults.isDefault(clear)) {
        oprot.writeFieldBegin(CLEAR_FIELD_DESC);
        oprot.writeBool(this.clear);
        oprot.writeFieldEnd();
      };
      java.util.Objects.requireNonNull(prepend, "prepend must not be null");
      
      if (!com.facebook.thrift.util.IntrinsicDefaults.isDefault(prepend)) {
        oprot.writeFieldBegin(PREPEND_FIELD_DESC);
        List<Integer> _iter0 = prepend;
        oprot.writeListBegin(new TList(TType.I32, _iter0.size()));
        for (int _iter1 : _iter0) {
          oprot.writeI32(_iter1);
        }
        oprot.writeListEnd();
        oprot.writeFieldEnd();
      }
      java.util.Objects.requireNonNull(append, "append must not be null");
      
      if (!com.facebook.thrift.util.IntrinsicDefaults.isDefault(append)) {
        oprot.writeFieldBegin(APPEND_FIELD_DESC);
        List<Integer> _iter0 = append;
        oprot.writeListBegin(new TList(TType.I32, _iter0.size()));
        for (int _iter1 : _iter0) {
          oprot.writeI32(_iter1);
        }
        oprot.writeListEnd();
        oprot.writeFieldEnd();
      }
      oprot.writeFieldStop();
      oprot.writeStructEnd();
    }
    
    private static class _RefFieldsField5PatchLazy {
        private static final RefFieldsField5Patch _DEFAULT = new RefFieldsField5Patch.Builder().build();
    }
    
    public static RefFieldsField5Patch defaultInstance() {
        return  _RefFieldsField5PatchLazy._DEFAULT;
    }
}
