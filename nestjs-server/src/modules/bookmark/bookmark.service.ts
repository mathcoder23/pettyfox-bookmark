import {Injectable} from '@nestjs/common';
import {Bookmark} from "./bookmark.entity";
import {SonicService} from "../sonic/sonic.service";

let md5 = require("md5")
import {RedisService} from "@liaoliaots/nestjs-redis";

@Injectable()
export class BookmarkService {
    private redis: any
    static KEY = "bookmark"

    constructor(private readonly sonicService: SonicService, private redisService: RedisService) {
        this.getRedis()
    }

    private async getRedis() {
        this.redis = await this.redisService.getClient()
    }

    add(bookmark: Bookmark): void {
        console.log('b', bookmark)
        bookmark.id = md5(bookmark.url)
        this.redis.hset(BookmarkService.KEY, bookmark.id, JSON.stringify(bookmark))
    }

    async list(): Promise<Bookmark[]> {
        let list = await this.redis.hvals(BookmarkService.KEY)
        let temp = []
        for (let item of list) {
            temp.push(JSON.parse(item))
        }
        console.log('list', temp)
        return temp
    }

    remove(ids: string[]): void {

    }

    update(bookmark: Bookmark): void {

    }
}
