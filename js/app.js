const nsq = require('nsqjs')

const w = new nsq.Writer('127.0.0.1', 4150)

w.connect()

w.on('ready', () => {
    w.deferPublish('topic_string', ['it really tied the room together'], 10000)
    w.deferPublish('topic_string', ['This message gonna arrive 5 sec later.'], 5000)
    w.publish('topic_string', [
        'Uh, excuse me. Mark it zero. Next frame.',
        'Smokey, this is not \'Nam. This is bowling. There are rules.'
    ])
    w.publish('topic_string', 'Wu?', err => {
        if (err) {
            return console.error(err.message)
        }
        console.log('Message sent successfully')
        w.close()
    })

    var _json = {
        "func": "evoke",
        "params": [
            "a",
            "b",
            "c"
        ]
    }

    w.publish('topic_json', _json, err => {
        if (err) {
            return console.error(err.message)
        }
        console.log('Message JSON sent successfully')
        w.close()
    })
})

w.on('closed', () => {
    console.log('Writer closed')
})