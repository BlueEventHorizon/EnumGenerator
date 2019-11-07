//
//  AppResource.swift
//  EnumGeneratorTest
//
//  Created by 寺田 克彦 on 2019/11/07.
//  Copyright © 2019 Katsuhiko Terada. All rights reserved.
//

import UIKit

let R = AppResource()

struct AppResource {
    let image = ImageResource()
    let color = ColorResource()
    let string = StringResource()
}

extension AppResource {
    struct ImageResource {
        let email_bcc = #imageLiteral(resourceName: "email_bcc")
    }
}

extension AppResource {
    struct ColorResource {
        
    }
}

extension AppResource {
    struct StringResource {
        
    }
}

class AppResourceTest {
    func test() {
        let _ = R.image.email_bcc
    }
}
