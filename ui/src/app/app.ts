import { Component, signal } from '@angular/core';
// import { RouterOutlet } from '@angular/router';
import { CurrencyPipe, DatePipe, TitleCasePipe, UpperCasePipe } from '@angular/common';

@Component({
  selector: 'app-root',
  imports: [UpperCasePipe],
  templateUrl: './app.html',
  styleUrl: './app.css'
})
export class App {
  protected readonly title = signal('ui');
  names: string[] = ["Samuel", "Gabriel", "Neuza", "Alexandre"]
}
