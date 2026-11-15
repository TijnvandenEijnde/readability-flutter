// swift-tools-version: 5.9
// The swift-tools-version declares the minimum version of Swift required to build this package.

import PackageDescription

let package = Package(
    name: "readability",
    platforms: [
        .iOS("13.0")
    ],
    products: [
        .library(name: "readability", targets: ["readability"])
    ],
    dependencies: [
        .package(name: "FlutterFramework", path: "../FlutterFramework")
    ],
    targets: [
        // Static Go c-archive built by src/Makefile (`make ios`).
        .binaryTarget(
            name: "libreadability",
            path: "libreadability.xcframework"
        ),
        // Small C target that references the Go exports so the linker keeps them.
        .target(
            name: "readability",
            dependencies: [
                "libreadability",
                .product(name: "FlutterFramework", package: "FlutterFramework"),
            ],
            linkerSettings: [
                .linkedFramework("CoreFoundation"),
                .linkedFramework("Security"),
                .linkedLibrary("resolv"),
            ]
        ),
    ]
)
