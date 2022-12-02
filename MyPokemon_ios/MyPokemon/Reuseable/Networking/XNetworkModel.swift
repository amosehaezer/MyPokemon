//
//  XNetworkModel.swift
//  MyPokemon
//
//  Created by Amos Ebenhaezer on 27/11/22.
//

import Foundation
import Moya

private func JSONResponseDataFormatter(_ data: Data) -> String {
    do {
        let dataAsJSON = try JSONSerialization.jsonObject(with: data)
        let prettyData = try JSONSerialization.data(withJSONObject: dataAsJSON, options: .prettyPrinted)
        return String(data: prettyData, encoding: .utf8) ?? String(data: data, encoding: .utf8) ?? ""
    } catch {
        return String(data: data, encoding: .utf8) ?? ""
    }
}

public class XNetworkModel {
    
    let provider = MoyaProvider<MainMoyaTarget>(plugins: [NetworkLoggerPlugin(
        configuration: .init(
            formatter: .init(responseData: JSONResponseDataFormatter)
            //logOptions: .verbose
        )
    )])
    
    public static let shared = XNetworkModel()
    typealias JSONStandard = [String: AnyObject]
    
    func fetchPokemon(sender: UIViewController,_ completion: @escaping (_ data: WrapperPokemon?) ->Void) {
        var data : WrapperPokemon!
        provider.request(.getListPokemon) { (result) in
           switch result{
           case .success(let response):
               print(response)
               let dataRespon = try! JSONDecoder().decode(WrapperPokemon.self,from: response.data)
               data = dataRespon
               completion(data)
           case .failure(let error):
               print(error.errorDescription)
           }
       }
    }
    
    
}
