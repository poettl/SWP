import { Component } from '@angular/core';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css'],
})
export class AppComponent {
  itemsSelected = 0;
  selectedItems = [];

  listSelectionChanged(event: any) {
    this.itemsSelected = this.selectedItems?.length;
  }
}
