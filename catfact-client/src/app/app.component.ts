import { Component } from '@angular/core';

import { CatFactServiceClient } from '../proto/catfact_grpc_web_pb';
import { CatFactRequest, CatFactResponse } from '../proto/catfact_pb';

@Component({
  selector: 'app-root',
  template: `
    <div id="factPlaceholder">Fact Goes Here</div>
    <button type="button" class="btn btn-primary">Get Fact</button>
  `
})
export class AppComponent {

  public ngOnInit(): void {

  }

}
