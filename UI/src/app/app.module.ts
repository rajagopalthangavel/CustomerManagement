import { Injectable, NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { LoginComponent } from './login/login.component';
import { MatButtonModule} from '@angular/material/button';
import { MatCardModule} from '@angular/material/card';
import { MatInputModule} from '@angular/material/input';
import { MatFormFieldModule} from '@angular/material/form-field';
import { MatTooltipModule} from '@angular/material/tooltip';
import { MatToolbarModule} from '@angular/material/toolbar';
import { MatIconModule } from '@angular/material/icon';
import { HttpClientModule, HTTP_INTERCEPTORS } from '@angular/common/http';
import { ToastrModule } from 'ngx-toastr';
import { FormsModule } from '@angular/forms';
import { InterInterceptor } from './inter.interceptor';
import {MatMenuModule} from '@angular/material/menu';
import { DatePipe } from '@angular/common';

@Injectable({
  providedIn: 'root'
})

@NgModule({
  declarations: [
    AppComponent,
    LoginComponent,
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    BrowserAnimationsModule,MatButtonModule,MatCardModule,MatInputModule,MatFormFieldModule,MatTooltipModule,MatToolbarModule,FormsModule,MatIconModule,
    HttpClientModule,ToastrModule.forRoot({timeOut:1500}),MatMenuModule
  ],
  providers: [{provide:HTTP_INTERCEPTORS,useClass:InterInterceptor,multi:true},DatePipe],
  bootstrap: [AppComponent]
})
export class AppModule { }
