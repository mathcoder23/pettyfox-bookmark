import {Module} from '@nestjs/common';
import {BookmarkService} from './bookmark.service';
import {BookmarkController} from "./bookmark.controller";
import {SonicModule} from "../sonic/sonic.module";
import {RedisModule} from "@liaoliaots/nestjs-redis";


@Module({
    imports: [SonicModule, RedisModule.forRoot({
        config: {
            host: 'localhost',
            port: 6379,
            db: 1

            // or with URL
            // url: 'redis://:your_password@localhost:6380/0'
        }
    })],
    providers: [BookmarkService],
    controllers: [BookmarkController]
})
export class BookmarkModule {
}
