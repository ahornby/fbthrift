/**
 * Autogenerated by Thrift
 *
 * DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING
 *  @generated
 */

package ;

import java.util.List;
import java.util.Map;
import java.util.Set;
import javax.annotation.concurrent.Immutable;
import javax.annotation.Nullable;
import com.facebook.hyperthrift.HyperThriftBase;
import com.facebook.hyperthrift.reflect.HyperThriftType;

@Immutable
@HyperThriftType
public class CrazyNesting extends HyperThriftBase {
  public static final String TYPE_NAME = "thrift.test.CrazyNesting";


  @Nullable
  public String string_field() {
    return getFieldValue(0);
  }

  @Nullable
  public Set<thrift.test.Insanity> set_field() {
    return getFieldValue(1);
  }

  public List<Map<Set<Integer>, Map<Integer, Set<List<Map<thrift.test.Insanity, String>>>>>> list_field() {
    return getFieldValue(2);
  }

  @Nullable
  public byte[] binary_field() {
    return getFieldValue(3);
  }

  @Nullable
  public List<Integer> i32_list_field() {
    return getFieldValue(4);
  }

  @Nullable
  public Set<Long> i64_set_field() {
    return getFieldValue(5);
  }



  public static class Builder extends HyperThriftBase.Builder {
    public Builder() {
      super(6);
    }

    public Builder(CrazyNesting other) {
      super(other);
    }

    @Nullable
    public String string_field() {
      return getFieldValue(0);
    }

    public Builder CrazyNesting( String value) {
      setFieldValue(0, value);
      return this;
    }

    @Nullable
    public Set<thrift.test.Insanity> set_field() {
      return getFieldValue(1);
    }

    public Builder CrazyNesting( Set<thrift.test.Insanity> value) {
      setFieldValue(1, value);
      return this;
    }

    @Nullable
    public List<Map<Set<Integer>, Map<Integer, Set<List<Map<thrift.test.Insanity, String>>>>>> list_field() {
      return getFieldValue(2);
    }

    public Builder CrazyNesting(@Nullable List<Map<Set<Integer>, Map<Integer, Set<List<Map<thrift.test.Insanity, String>>>>>> value) {
      setFieldValue(2, value);
      return this;
    }

    @Nullable
    public byte[] binary_field() {
      return getFieldValue(3);
    }

    public Builder CrazyNesting( byte[] value) {
      setFieldValue(3, value);
      return this;
    }

    @Nullable
    public List<Integer> i32_list_field() {
      return getFieldValue(4);
    }

    public Builder CrazyNesting( List<Integer> value) {
      setFieldValue(4, value);
      return this;
    }

    @Nullable
    public Set<Long> i64_set_field() {
      return getFieldValue(5);
    }

    public Builder CrazyNesting( Set<Long> value) {
      setFieldValue(5, value);
      return this;
    }

    public CrazyNesting build() {
      Object[] fields = markBuilt();
      CrazyNesting instance = new CrazyNesting();
      instance.init(TYPE_NAME, fields);
      instance.assertRequired(2, "list_field");
      return instance;
    }
  }
}