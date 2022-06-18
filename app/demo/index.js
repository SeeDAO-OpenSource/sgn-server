
var count = 0
var limit = 20
window.onload = (h, e) => {
    load()
}

const baseURL = ""

const load = async () => {
    const sgnList = document.getElementById("sgnList");
    const response = await fetch(`${baseURL}/api/sgn?skip=${count}&limit=${limit}`)
    const data = await response.json()
    if (data.success && data.data) {
        data.data.forEach(ntf => {
            sgnList.appendChild(createSgnItem(ntf))
        });
        count += data.data.length
    }
}

const createSgnItem = (data) => {
    const item = document.createElement('div')
    item.className = 'sgn-item'
    const name = document.createElement('div')
    name.className = 'sgn-item-name'
    name.innerText = data.token_id
    const img = document.createElement('img')
    img.className = 'sgn-item-img'
    img.src = `${baseURL}/api/sgn/image/${data.token_id}?w=120&h=120`
    item.appendChild(name)
    item.appendChild(img)
    return item
}

const loadMore = async () => {
    await load()
}

const setLoadLimit = async (size) => {
    if (limit != size) {
        limit = size
        await load();
    }
}

const submitPageSize = async () => {
    const size = document.getElementById("pagesizeInput").value
    if (size) {
        await setLoadLimit(size)
    }
}