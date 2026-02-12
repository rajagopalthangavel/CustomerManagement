import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';

import { UserRoutingModule } from './user-routing.module';
import { ListUserComponent } from './list-user/list-user.component';
import {MatToolbarModule} from '@angular/material/toolbar';
import {MatDividerModule} from '@angular/material/divider';
import { MatTableModule } from '@angular/material/table'  
import {MatIconModule} from '@angular/material/icon';
import {MatTooltipModule} from '@angular/material/tooltip';
import {MatPaginatorModule} from '@angular/material/paginator';
import { HttpClientModule, HTTP_INTERCEPTORS } from '@angular/common/http'
import { MatButtonModule } from '@angular/material/button';
import { AddUserComponent } from './add-user/add-user.component';
import {MatCardModule} from '@angular/material/card';
import {MatFormFieldModule} from '@angular/material/form-field';
import {MatInputModule} from '@angular/material/input';
import { FormsModule } from '@angular/forms';
import { UpdateUserComponent } from './update-user/update-user.component';
import {MatSelectModule} from '@angular/material/select';
import { MatOptionModule } from '@angular/material/core';
import {MatDialogModule} from '@angular/material/dialog';
import { DeleteDialogComponent } from './delete-dialog/delete-dialog.component';
import {MatMenuModule} from '@angular/material/menu';
import { InterInterceptor } from '../inter.interceptor';

@NgModule({
  declarations: [
    ListUserComponent,
    AddUserComponent,
    UpdateUserComponent,
    DeleteDialogComponent,
  ],
  imports: [
    CommonModule,
    UserRoutingModule,MatToolbarModule,MatDividerModule,MatTableModule,MatIconModule,MatTooltipModule,MatPaginatorModule,HttpClientModule,MatButtonModule,
    MatCardModule,MatFormFieldModule,MatInputModule,FormsModule,MatSelectModule,MatOptionModule,MatDialogModule,MatMenuModule
  ],
  providers: [{provide:HTTP_INTERCEPTORS,useClass:InterInterceptor,multi:true}]
})
export class UserModule { }
