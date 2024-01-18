// This is an empty Ballerina class autogenerated to represent the `java.util.Set` Java interface.
//
// If you need the implementation of this class generated, please use the following command.
//
// $ bal bindgen [(-cp|--classpath) <classpath>...]
//               [(-mvn|--maven) <groupId>:<artifactId>:<version>]
//               [(-o|--output) <output-path>]
//               [--public]
//               (<class-name>...)
//
// E.g. $ bal bindgen java.util.Set

import ballerina/jballerina.java;

# Ballerina class mapping for the Java `java.util.Set` interface.
@java:Binding {'class: "java.util.Set"}
public distinct class Set {

    *java:JObject;

    # The `handle` field that stores the reference to the `java.util.Set` object.
    public handle jObj;

    # The init function of the Ballerina class mapping the `java.util.Set` Java interface.
    #
    # + obj - The `handle` value containing the Java reference of the object.
    public function init(handle obj) {
        self.jObj = obj;
    }

    # The function to retrieve the string representation of the Ballerina class mapping the `java.util.Set` Java interface.
    #
    # + return - The `string` form of the Java object instance.
    public function toString() returns string {
        return java:toString(self.jObj) ?: "";
    }

}

