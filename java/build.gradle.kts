plugins {
    java
    id("com.autonomousapps.dependency-analysis") version "1.28.0"
    id("com.diffplug.spotless") version "6.23.3"
    id("com.github.ben-manes.versions") version "0.50.0"
    id("com.google.protobuf") version "0.9.4"
    id("se.ascp.gradle.gradle-versions-filter") version "0.1.16"
}

sourceSets {
    main {
        java {
            srcDir("build/generated/source/protoc-gen-twirp-java/main/java")
        }
        proto {
            srcDir("../protobuf")
        }
    }
}

dependencies {
    implementation(platform("io.helidon:helidon-bom:2.6.5"))

    implementation("com.google.protobuf:protobuf-java:3.25.1")
    implementation("io.helidon.webserver:helidon-webserver")
}

repositories {
    mavenCentral()
}

protobuf {
    protoc {
        artifact = "com.google.protobuf:protoc:3.25.1"
    }
}

versionsFilter {
    gradleReleaseChannel.set("current")
    checkConstraints.set(true)
    outPutFormatter.set("json")
}

spotless {
    java {
        googleJavaFormat("1.18.1").reflowLongStrings().skipJavadocFormatting()
        formatAnnotations()
        targetExclude("build/**")
    }
}
