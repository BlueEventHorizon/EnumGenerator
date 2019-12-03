//
//  AppResource.swift
//  EnumGeneratorTest
//
//  Created by k2moons on 2019/11/07.
//  Copyright © 2019 k2moons. All rights reserved.
//

// https://developer.apple.com/library/archive/documentation/Xcode/Reference/xcode_ref-Asset_Catalog_Format/index.html

import Foundation

let R = AppResource()

struct AppResource {
    var Image: ImageResource.Type { return ImageResource.self }
    var Color: ColorResource.Type { return ColorResource.self }
    var String: StringResource.Type { return StringResource.self }
}

extension AppResource {
    struct ImageResource {}
}

extension AppResource {
    struct ColorResource {}
}

extension AppResource {
    struct StringResource {}
}

// MARK: - Test

class AppResourceTest {

}
