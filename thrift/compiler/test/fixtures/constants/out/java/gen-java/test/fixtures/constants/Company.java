/**
 * Autogenerated by Thrift
 *
 * DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING
 *  @generated
 */

package test.fixtures.constants;

import com.facebook.swift.codec.*;
import org.apache.thrift.TException;
import org.apache.thrift.protocol.TProtocol;

@SwiftGenerated
public enum Company implements com.facebook.thrift.enums.BaseEnum, com.facebook.thrift.payload.ThriftSerializable {
    FACEBOOK(0),
    WHATSAPP(1),
    OCULUS(2),
    INSTAGRAM(3),
    __FRIEND__FEED(4);

    private final int value;

    Company(int value) {
        this.value = value;
    }

    @ThriftEnumValue
    public int getValue() {
        return value;
    }

    public static Company fromInteger(int n) {
        switch (n) {
        case 0:
            return FACEBOOK;
        case 1:
            return WHATSAPP;
        case 2:
            return OCULUS;
        case 3:
            return INSTAGRAM;
        case 4:
            return __FRIEND__FEED;
        default:
            return null;
        }
    }

    public static com.facebook.thrift.payload.Reader<Company> asReader() {
        return Company::read0;
    }

    public static Company read0(TProtocol iprot) throws TException {
        return Company.fromInteger(iprot.readI32());
    }

    public void write0(TProtocol oprot) throws TException {
        oprot.writeI32(getValue());
    }


}
