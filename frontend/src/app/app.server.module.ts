import { NgModule } from '@angular/core';
import { ServerModule } from '@angular/platform-server';

import { AppModule } from './app.module';
import { AppComponent } from './app.component';
import { InlineStyleModule } from './inline-style/inline-style.module';
import { InlineStyleComponent } from './inline-style/inline-style.component';

@NgModule({
  imports: [AppModule, ServerModule, InlineStyleModule],
  bootstrap: [AppComponent, InlineStyleComponent],
})
export class AppServerModule {}
