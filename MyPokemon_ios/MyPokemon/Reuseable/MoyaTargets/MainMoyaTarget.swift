//
//  MainMoyaTarget.swift
//  MyPokemon
//
//  Created by Amos Ebenhaezer on 27/11/22.
//

import Foundation
import Moya

enum MainMoyaTarget {
    case getListPokemon
}

extension MainMoyaTarget: TargetType {
    var method: Moya.Method {
        switch self {
            case .getListPokemon:
                return .get
        }
    }
    
    
    var task: Moya.Task {
        switch self {
        case .getListPokemon:
            return .requestPlain
        }
    }
    
    var headers: [String : String]? {
        return ["Content-type": "application/json"]
    }
    
    
    var baseURL: URL {
        return URL(string: XhttpConfig.AppDomain())!
    }
    
    var path: String {
        switch self {
        case .getListPokemon:
            return "/pokemon/list"
        }
    }
    
    var sampleData: Data {
        switch self {
        case .getListPokemon:
            return XJsonLoadFileFactory().from(fileName: "pokemon")
        }
    }
}
