const faker = require('faker')
const fs = require('fs')
const async = require('neo-async')
const moment = require('moment')
const argv = process.argv
const functionToExec = argv[2]
const numOfItems = argv[3]

switch (functionToExec) {
  case 'widgets':
    genWidgets(numOfItems, err => {
      if (!err) console.log('Done Widgets')
    })
    break
  case 'users':
    genUsers(numOfItems, err => {
      if (!err) console.log('Done Users')
    })
    break
  default: displayUsage()
}

function getUserLine (name, gravatarId) {
  const date = moment().format('YYYY-MM-DD')
  return `('${name}', 'http://www.gravatar.com/avatar/${gravatarId}?s=64&d=monsterid', '${date}'),`
}

function getWidgetLine (name, color, price, inventory, melts) {
  const date = moment().format('YYYY-MM-DD')
  return `('${name}', '${color}', '${price}', ${inventory}, ${melts}, '${date}'),`
}

function genUsers (numOfUsers, cb) {
  let users = []
  for (var i = 0; i < numOfUsers; i++) {
    const name = faker.name.findName()
    const gravatarId = faker.random.alphaNumeric(32)
    const line = getUserLine(name, gravatarId)
    users.push(line)
  }
  async.each(users, (user, cbe) => {
    fs.appendFile('userData.txt', `${user}\n`, cbe)
  }, cb)
}

function genWidgets (numOfWidgets, cb) {
  let widgets = []
  for (var i = 0; i < numOfWidgets; i++) {
    const name = faker.commerce.productName()
    const color = faker.commerce.color()
    const price = faker.commerce.price()
    const inventory = faker.random.number({min: 0, max: 1000})
    const melts = faker.random.boolean()

    const line = getWidgetLine(name, color, price, inventory, melts)
    widgets.push(line)
  }
  async.each(widgets, (widget, cbe) => {
    fs.appendFile('widgetData.txt', `${widget}\n`, cbe)
  }, cb)
}

function displayUsage () {
  console.log(`
Usage: node generate.js <function> <numOfItems>
  `)
}
