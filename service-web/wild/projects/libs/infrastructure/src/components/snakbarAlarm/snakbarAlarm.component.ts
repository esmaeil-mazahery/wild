import { AlertService } from './../../serivces/system/alert.service';
import { Component, OnInit } from '@angular/core';
import { MatSnackBar } from '@angular/material/snack-bar';

@Component({
  selector: 'app-snakbar-alarm',
  templateUrl: './snakbarAlarm.component.html',
  styleUrls: ['./snakbarAlarm.component.scss'],
})
export class SnakbarAlarmComponent implements OnInit {
  constructor(AlertService: AlertService, private snackBar: MatSnackBar) {
    AlertService.getMessage().subscribe((msg) => {
      if (msg && msg.type == 'SnackBar')
        this.snackBar.open(msg.text, '', {
          duration: msg.data!.duration,
        });
    });
  }

  ngOnInit() {}
}
