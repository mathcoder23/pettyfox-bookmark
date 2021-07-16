export class ObjResponse {
    code: number
    data: object
    msg: string

    static ok(data): ObjResponse {
        let obj = new ObjResponse()
        obj.code = 200
        obj.data = data
        return obj
    }

    static error(msg): ObjResponse {
        let obj = new ObjResponse()
        obj.code = 500
        obj.msg = msg
        return obj
    }
}
