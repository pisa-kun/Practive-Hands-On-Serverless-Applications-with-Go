import { NgModule, CUSTOM_ELEMENTS_SCHEMA } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
// REST クライアント実装ののためのサービスを import ( Angular 5.0.0 以降はこちらを使う )
import { HttpClientModule } from '@angular/common/http';
import { NgbModule } from '@ng-bootstrap/ng-bootstrap';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { NavbarComponent } from './components/navbar/navbar.component';
import { IdolItemComponent } from './components/idol-item/idol-item.component';
import { ListIdolsComponent } from './components/list-idols/list-idols.component';
import { IdolApiService } from './services/idol-api.service';
import { NewIdolComponent } from './components/new-idol/new-idol.component';
import { CognitoService } from './services/cognito.service';

import { StorageServiceModule} from 'ngx-webstorage-service';

@NgModule({
  declarations: [
    AppComponent,
    NavbarComponent,
    IdolItemComponent,
    ListIdolsComponent,
    NewIdolComponent
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    HttpClientModule,
    StorageServiceModule
  ],
  providers: [
    IdolApiService,
    CognitoService
  ],
  bootstrap: [AppComponent],
  schemas: [ CUSTOM_ELEMENTS_SCHEMA ]
})
export class AppModule { }
