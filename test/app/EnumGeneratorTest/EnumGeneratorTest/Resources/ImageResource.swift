//
//  ImageResource.swift
//  EnumGeneratorTest
//
//  Created by 寺田 克彦 on 2019/11/07.
//  Copyright © 2019 Katsuhiko Terada. All rights reserved.
//

import UIKit.UIImage

extension AppResource.ImageResource {
    struct Communication {
        static var email_bcc: UIImage { return #imageLiteral(resourceName: "email_bcc") }
    }
}

// MARK: - Test

extension AppResourceTest {
    func testImageResource() {
        let _ = AppResource.ImageResource.Communication.email_bcc
        let _ = R.Image.Communication.email_bcc
    }
}
