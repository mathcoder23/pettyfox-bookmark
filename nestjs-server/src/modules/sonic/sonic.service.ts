import {Injectable} from '@nestjs/common';
import {Ingest, Search} from "sonic-channel/lib/sonic_channel";

let sonicHost = process.env.SONIC_HOST
@Injectable()
export class SonicService {
    search = null
    ingest = null

    getSearch(): Promise<Search> {
        return new Promise<Search>(((resolve, reject) => {
            if (this.search != null) {
                resolve(this.search)
                return
            }
            let tempSearch = new Search({
                host: sonicHost,            // Or '127.0.0.1' if you are still using IPv4
                port: 1491,             // Default port is '1491'
                auth: "SecretPassword"  // Authentication password (if any)
            }).connect({
                connected: () => {
                    this.search = tempSearch
                    resolve(this.search)
                    // Connected handler
                    console.info("Sonic Channel succeeded to connect to host (search).");
                },

                disconnected: () => {
                    // Disconnected handler
                    console.error("Sonic Channel is now disconnected (search).");
                    this.search = null
                    reject()
                },

                timeout: () => {
                    // Timeout handler
                    console.error("Sonic Channel connection timed out (search).");
                    this.search = null

                    reject()

                },

                retrying: () => {
                    // Retry handler
                    console.error("Trying to reconnect to Sonic Channel (search)...");
                    this.search = null

                    reject()

                },
                error: (error) => {
                    // Failure handler
                    console.error("Sonic Channel failed to connect to host (search).", error);
                    this.search = null

                    reject()

                }
            });
        }))

    }

    getIngest(): Promise<Ingest> {
        return new Promise<Ingest>(((resolve, reject) => {
            if (this.ingest != null) {
                resolve(this.ingest)
                return
            }
            let tempSearch = new Ingest({
                host: sonicHost,            // Or '127.0.0.1' if you are still using IPv4
                port: 1491,             // Default port is '1491'
                auth: "SecretPassword"  // Authentication password (if any)
            }).connect({
                connected: () => {
                    this.ingest = tempSearch
                    resolve(this.ingest)
                    // Connected handler
                    console.info("Sonic Channel succeeded to connect to host (search).");
                },

                disconnected: () => {
                    // Disconnected handler
                    console.error("Sonic Channel is now disconnected (search).");
                    this.ingest = null
                    reject()
                },

                timeout: () => {
                    // Timeout handler
                    console.error("Sonic Channel connection timed out (search).");
                    this.ingest = null

                    reject()

                },

                retrying: () => {
                    // Retry handler
                    console.error("Trying to reconnect to Sonic Channel (search)...");
                    this.ingest = null

                    reject()

                },
                error: (error) => {
                    // Failure handler
                    console.error("Sonic Channel failed to connect to host (search).", error);
                    this.ingest = null

                    reject()

                }
            });
        }))

    }

    async add(): Promise<void> {
        let search = await this.getSearch()
        let ingest = await this.getIngest()
        await ingest.push("message", "default", "123", "测试你好")
        let results = await search.query("message", "default", "测试")
        console.log('aaaaa', results)
    }

    async searchData(collection, bucket,  text): Promise<string[]> {
        let search = await this.getSearch()
        let results = await search.query(collection, bucket, text)
        return results
    }

    addData(collection, bucket, object, text): void {
        this.getIngest().then((ingest) => {
            ingest.push(collection, bucket, object, text).then(r => {
            })
        })
    }
}
