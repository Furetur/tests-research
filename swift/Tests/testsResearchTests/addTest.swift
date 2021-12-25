import XCTest
@testable import testsResearch

class AddTest : XCTestCase {
    func testAdd() {
        XCTAssertEqual(add(x: 5, y: 6), 12)
    }
}