import { Component, OnInit } from '@angular/core';
import { Mediator } from './mediator.service';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css'],
})
export class AppComponent implements OnInit {
  selectedItems = [];
  mediator = new Mediator();
  button1Disabled = false;
  button2Disabled = false;
  button3Disabled = false;
  constructor() {
    this.mediator.addObserver('listCount', () => {
      const value = this.mediator.getValue('listCount');
      if (value === 1) {
        this.button1Disabled = false;
        this.button2Disabled = true;
        this.button3Disabled = true;
      } else if (value >= 1) {
        this.button1Disabled = true;
        this.button2Disabled = false;
        this.button3Disabled = true;
      } else if (value === 0) {
        this.button1Disabled = true;
        this.button2Disabled = true;
        this.button3Disabled = false;
      }
    });
  }
  ngOnInit(): void {
    this.mediator.setValue('listCount', 0);
  }

  listSelectionChanged(event: any) {
    this.mediator.setValue('listCount', this.selectedItems?.length);
  }
}
