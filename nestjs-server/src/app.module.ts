import { Module } from '@nestjs/common';
import { AppController } from './app.controller';
import { AppService } from './app.service';
import {BookmarkModule} from "./modules/bookmark/bookmark.module";
import { SonicModule } from './modules/sonic/sonic.module';

@Module({
  imports: [BookmarkModule, SonicModule],
  controllers: [AppController],
  providers: [AppService],
})
export class AppModule {}
