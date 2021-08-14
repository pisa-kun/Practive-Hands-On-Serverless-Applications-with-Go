import { Injectable } from '@angular/core';
// import 'rxjs/add/operator/map';
import { environment } from '../../environments/environment';
// REST クライアント実装ののためのサービスを import ( Angular 5.0.0 以降はこちらを使う )
import { HttpClient } from '@angular/common/http';
import { Idol } from '../models/idol';

@Injectable()
export class IdolApiService {

  constructor(private http:HttpClient) { }

  public findAll(){
    console.log("call findAll in IolApiService");
    // return this.http
    //   .get(environment.api)
    //   .map(res => {
    //     return res.json()
    //   });
    return this.http.get(environment.api)
    .toPromise()
    .then((res) => {
      // response の型は any ではなく class で型を定義した方が良いが
      // ここでは簡便さから any としておく

      // @angular/http では json() でパースする必要があったが､ @angular/common/http では不要となった
      //const response: any = res.json();
      const response: any = res;
      console.log(response);
      return response;
    })
    .catch(this.errorHandler);
  }

  insert(idol: Idol){
    idol.addId("114");
    console.log("insert +", idol)
    return this.http
      .post(environment.api, JSON.stringify(idol))
      // .map(res => {
      //   return res
      // })
    }
    
    /**
   * REST-API 実行時のエラーハンドラ
   * (toPromise.then((res) =>{}) を利用する場合のコード)
   *
   * @private
   * @param {any} err エラー情報
   * @memberof HttpClientService
   */
     private errorHandler(err:any) {
      console.log('Error occured.', err);
      return Promise.reject(err.message || err);
    }
}