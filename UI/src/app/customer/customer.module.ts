import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';

import { CustomerRoutingModule } from './customer-routing.module';
import { ListCustomerComponent } from './list-customer/list-customer.component';
import {MatToolbarModule} from '@angular/material/toolbar';
import {MatButtonModule} from '@angular/material/button';
import {MatMenuModule} from '@angular/material/menu';
import {MatIconModule} from '@angular/material/icon';
import {MatTooltipModule} from '@angular/material/tooltip';
import {MatTableModule} from '@angular/material/table';
import {MatPaginatorModule} from '@angular/material/paginator';
import { AddCustomerComponent } from './add-customer/add-customer.component';
import {MatCardModule} from '@angular/material/card';
import {MatFormFieldModule} from '@angular/material/form-field';
import {MatInputModule} from '@angular/material/input';
import { FormsModule } from '@angular/forms';
import { UpdateCustomerComponent } from './update-customer/update-customer.component';
import {MatDialogModule} from '@angular/material/dialog';


@NgModule({
  declarations: [
    ListCustomerComponent,
    AddCustomerComponent,
    UpdateCustomerComponent
  ],
  imports: [
    CommonModule,
    CustomerRoutingModule,MatToolbarModule,MatButtonModule,MatMenuModule,MatIconModule,MatTooltipModule,MatTableModule,MatPaginatorModule,
    MatCardModule,MatInputModule,MatFormFieldModule,FormsModule,MatDialogModule
  ]
})
export class CustomerModule { }
