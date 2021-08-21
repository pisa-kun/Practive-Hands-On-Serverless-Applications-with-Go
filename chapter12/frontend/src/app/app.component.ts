import { Component, Inject, ViewChild, OnInit, ViewChildren } from '@angular/core';
import { NgbModal, NgbModalRef } from '@ng-bootstrap/ng-bootstrap';
import { CognitoService } from './services/cognito.service';
import {LOCAL_STORAGE, WebStorageService} from 'ngx-webstorage-service';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})

export class AppComponent implements OnInit{
  // private loginModal: NgbModalRef = {};
  loginModal: any = {};
  loginError : boolean = false;
  title = "frontend"

  @ViewChild('login',{static: true}) public content:any;

  constructor(private modalService: NgbModal,
              private cognitoService: CognitoService,
              @Inject(LOCAL_STORAGE) private storage: WebStorageService){
                console.log("create app.component")
              }

  ngOnInit(){
    if(!this.storage.get("COGNITO_TOKEN")){
      console.log("ngOnInit")
      console.log(this.content)
      this.loginModal = this.modalService.open(this.content)
    }
  }

  signin(username:any, password:any){
    console.log("signin")

    this.cognitoService.auth(username, password, (err:any, token:any) => {
      if(err){
        this.loginError = true
      }else{
        this.loginError = false
        this.storage.set("COGNITO_TOKEN", token)
        this.loginModal.close()
        window.location.reload()
      }
      console.log("sign in result", ":", err)
    })
  }
}
