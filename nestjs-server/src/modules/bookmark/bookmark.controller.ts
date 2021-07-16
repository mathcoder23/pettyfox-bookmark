import {Body, Controller, Get, Inject, Param, Post, Query} from '@nestjs/common';
import {Bookmark} from "./bookmark.entity";
import {BookmarkService} from "./bookmark.service";
import {ObjResponse} from "../../comms/api/obj-response";
import {IdsQuery} from "../../comms/api/ids-query";

@Controller('bookmark')
export class BookmarkController {
    constructor(@Inject(BookmarkService) private readonly bookmarkService: BookmarkService) {
    }


    @Get('search')
    search(@Query("keyword") keyword: string): ObjResponse {
        console.log('ke', keyword)
        return ObjResponse.ok(this.bookmarkService.search(keyword))
    }

    @Post('add')
    add(@Body() bookmark: Bookmark): ObjResponse {
        this.bookmarkService.add(bookmark)
        return ObjResponse.ok(null)
    }

    @Post('update')
    update(@Body() bookmark: Bookmark): ObjResponse {
        this.bookmarkService.update(bookmark)
        return ObjResponse.ok(null)
    }

    @Post('remove')
    remove(@Body() query: IdsQuery): ObjResponse {
        this.bookmarkService.remove(query.ids)
        return ObjResponse.ok(null)
    }

    @Get('list')
    async list(): Promise<ObjResponse> {
        let rep = await this.bookmarkService.list()
        return ObjResponse.ok(rep)
    }
}
