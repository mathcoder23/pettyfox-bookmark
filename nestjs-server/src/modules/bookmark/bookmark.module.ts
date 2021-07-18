import {Module} from '@nestjs/common';
import {BookmarkService} from './bookmark.service';
import {BookmarkController} from "./bookmark.controller";
import {SonicModule} from "../sonic/sonic.module";
import {RedisModule} from "@liaoliaots/nestjs-redis";

let redisHost = process.env.REDIS_HOST
let redisPort = process.env.REDIS_PORT

@Module({
    imports: [SonicModule, RedisModule.forRoot({
        config: {
            host: redisHost,
            port: parseInt(redisPort),
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
