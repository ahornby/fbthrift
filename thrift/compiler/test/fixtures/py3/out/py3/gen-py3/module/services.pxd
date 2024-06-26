#
# Autogenerated by Thrift for thrift/compiler/test/fixtures/py3/src/module.thrift
#
# DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING
#  @generated
#

from thrift.py3.server cimport ServiceInterface


cdef class SimpleServiceInterface(ServiceInterface):
    # these are to avoid weird Cython multiple inheritance issue
    cdef bint _for_cython_get_five
    cdef bint _for_cython_add_five
    cdef bint _for_cython_do_nothing
    cdef bint _for_cython_concat
    cdef bint _for_cython_get_value
    cdef bint _for_cython_negate
    cdef bint _for_cython_tiny
    cdef bint _for_cython_small
    cdef bint _for_cython_big
    cdef bint _for_cython_two
    cdef bint _for_cython_expected_exception
    cdef bint _for_cython_unexpected_exception
    cdef bint _for_cython_sum_i16_list
    cdef bint _for_cython_sum_i32_list
    cdef bint _for_cython_sum_i64_list
    cdef bint _for_cython_concat_many
    cdef bint _for_cython_count_structs
    cdef bint _for_cython_sum_set
    cdef bint _for_cython_contains_word
    cdef bint _for_cython_get_map_value
    cdef bint _for_cython_map_length
    cdef bint _for_cython_sum_map_values
    cdef bint _for_cython_complex_sum_i32
    cdef bint _for_cython_repeat_name
    cdef bint _for_cython_get_struct
    cdef bint _for_cython_fib
    cdef bint _for_cython_unique_words
    cdef bint _for_cython_words_count
    cdef bint _for_cython_set_enum
    cdef bint _for_cython_list_of_lists
    cdef bint _for_cython_word_character_frequency
    cdef bint _for_cython_list_of_sets
    cdef bint _for_cython_nested_map_argument
    cdef bint _for_cython_make_sentence
    cdef bint _for_cython_get_union
    cdef bint _for_cython_get_keys
    cdef bint _for_cython_lookup_double
    cdef bint _for_cython_retrieve_binary
    cdef bint _for_cython_contain_binary
    cdef bint _for_cython_contain_enum
    cdef bint _for_cython_get_binary_union_struct
    pass

cdef class DerivedServiceInterface(SimpleServiceInterface):
    # these are to avoid weird Cython multiple inheritance issue
    cdef bint _for_cython_get_six
    pass

cdef class RederivedServiceInterface(DerivedServiceInterface):
    # these are to avoid weird Cython multiple inheritance issue
    cdef bint _for_cython_get_seven
    pass

