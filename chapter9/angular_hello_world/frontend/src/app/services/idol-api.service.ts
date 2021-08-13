import { Injectable } from '@angular/core';
//import 'rxjs/add/operator/map';
import { environment } from '../../environments/environment';
// REST クライアント実装ののためのサービスを import ( Angular 5.0.0 以降はこちらを使う )
import { HttpClient, HttpHeaders } from '@angular/common/http';
//import { Idol } from '../idols/idol';

@Injectable()
export class IdolApiService {

  constructor(private http:HttpClient) { }

  findAll(){
    // return this.http
    //   .get(environment.api)
    //   .map(res => {
    //     return res.json()
    //   })
    return this.http.get(environment.api)
    .toPromise()
    .then((res) => {
      // response の型は any ではなく class で型を定義した方が良いが
      // ここでは簡便さから any としておく

      // @angular/http では json() でパースする必要があったが､ @angular/common/http では不要となった
      //const response: any = res.json();
      const response: any = res;
      return response;
    })
  }

}
