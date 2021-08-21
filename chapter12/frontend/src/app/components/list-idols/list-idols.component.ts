import { ThisReceiver } from '@angular/compiler';
import { Component, OnInit } from '@angular/core';
import { Idol } from '../../models/idol';
import { IdolApiService } from '../../services/idol-api.service';

@Component({
  selector: 'list-idols',
  templateUrl: './list-idols.component.html',
  styleUrls: ['./list-idols.component.css']
})
export class ListIdolsComponent implements OnInit {
/**
   * バックエンドから返却されたレスポンスをセットするプロパティ
   *
   * 型は any ではなく class で型を定義した方が良いが
   * ここでは簡便さから any としておく
   *
   * @private
   * @type {string}
   * @memberof HttpClientComponent
   */
 public param: any = {};

 /**
  * バックエンドから返却されたたメッセージをセットするプロパティ
  *
  * @type {*}
  * @memberof HttpClientComponent
  */
 public messageInfo: any = {
   body: null,
   statusCode: null,
 };

 /**
  * バックエンドから返却されたたメッセージを保持するリストプロパティ
  *
  * @type {*}
  * @memberof HttpClientComponent
  */
 public messageInfoList: any = [this.messageInfo];

 /**
  * メッセージ登録回数
  *
  * @private
  * @type {number}
  * @memberof HttpClientComponent
  */
 public messageId: number = 1;

 /**
  * 入力メッセージ
  *
  * @type {string}
  * @memberof HttpClientComponent
  */
 public message: string = '';

  public idols: Idol[];
  constructor(private idolApiService:IdolApiService) {
    console.log("constructor idolApiService");
    this.idols = []
    this.idolApiService.findAll()
    .then(
      (response) => {
        this.param = response;
        // bodyの一覧を格納
        this.messageInfoList = this.param.body;
        console.log(this.messageInfoList);
        const obj = JSON.parse(this.messageInfoList);
        console.log(obj);
        obj.forEach( (idol:any) => {
          this.idols.push(new Idol(idol.name, idol.description, idol.cover))
        });
        // this.param = response;
        // this.messageInfoList = this.param.messages;
      }
    )
    .catch(
      (error:any) => console.log(error)
    );
    // (res =>{
    //   res.array.forEach(idol => {
    //     this.idols.push(new Idol(idol.name, "some description"))
    //   });
    // })
   }

  ngOnInit(): void {
  }

}
