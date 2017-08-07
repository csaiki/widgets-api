const faker = require('faker')
const fs = require('fs')
const async = require('neo-async')

function getUserLine (index, name, gravatarId) {
  return `(${index}, '${name}', 'http://www.gravatar.com/avatar/${gravatarId}?s=64&d=monsterid')`
}

function getWidgetLine (index, name, color, price, inventory, melts) {
  return `(${index}, '${name}', '${color}', '${price}', ${inventory}, ${melts})`
}

function genUsers (numOfUsers, cb) {
  let users = []
  for (var i = 0; i < numOfUsers; i++) {
    const index = i
    const name = faker.name.findName()
    const gravatarId = faker.random.alphaNumeric(32)
    const line = getUserLine(index, name, gravatarId)
    users.push(line)
  }
  async.each(users, (user, cbe) => {
    fs.appendFile('userData.txt', `${user}\n`, cbe)
  }, cb)
}

function genWidgets (numOfWidgets, cb) {
  let widgets = []
  for (var i = 0; i < numOfWidgets; i++) {
    const index = i
    const name = faker.commerce.productName()
    const color = faker.commerce.color()
    const price = faker.commerce.price()
    const inventory = faker.random.number({min: 0, max: 1000})
    const melts = faker.random.boolean()

    const line = getWidgetLine(index, name, color, price, inventory, melts)
    widgets.push(line)
  }
  async.each(widgets, (widget, cbe) => {
    fs.appendFile('widgetData.txt', `${widget}\n`, cbe)
  }, cb)
}

genWidgets(100, (err, data) => {
  if (!err) console.log('done')
})
