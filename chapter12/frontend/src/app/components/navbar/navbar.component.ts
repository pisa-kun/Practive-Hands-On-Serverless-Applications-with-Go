import { Component, Inject } from '@angular/core';
import { NgbModal } from '@ng-bootstrap/ng-bootstrap';
import {LOCAL_STORAGE, WebStorageService} from 'ngx-webstorage-service';

@Component({
  selector: 'idols-navbar',
  templateUrl: './navbar.component.html',
  styleUrls: ['./navbar.component.css']
})
export class NavbarComponent{
  constructor(private modalService: NgbModal,
    @Inject(LOCAL_STORAGE) private storage: WebStorageService) {}

  openNewIdolModal(content:any){
  console.log(content)
  this.modalService.open(content)
  }

  signout(){
  this.storage.remove("COGNITO_TOKEN")
  window.location.reload()
  }

  isLogged(){
  return this.storage.get("COGNITO_TOKEN") ? true : false
  }
}