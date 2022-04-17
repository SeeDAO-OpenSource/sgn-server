
var page = 1
var pageSize = 20
window.onload = (h, e) => {
    init()
}

const baseURL = "http://localhost:5000"

const init = async () => {
    const nftList = document.getElementById("ntfList");
    chidren = [...nftList.childNodes]
    chidren.forEach(n => nftList.removeChild(n))
    const response = await fetch(`${baseURL}/api/v1/nft/0x23fDA8a873e9E46Dbe51c78754dddccFbC41CFE1?page=${page}&page_size=${pageSize}`)
    const data = await response.json()
    if (data.success && data.data) {
        data.data.forEach(ntf => {
            nftList.appendChild(createNtfItem(ntf))
        });
    }
}

const createNtfItem = (data) => {
    const item = document.createElement('div')
    item.className = 'nft-item'
    const name = document.createElement('div')
    name.className = 'nft-item-name'
    name.innerText = data.token_id
    const img = document.createElement('img')
    img.className = 'nft-item-img'
    img.src = `${baseURL}/api/v1/blob/${encodeURIComponent(data.metadata.image)}?w=120&h=120`
    item.appendChild(name)
    item.appendChild(img)
    return item
}

const prePage = async () => {
    page--
    if (page < 1) {
        page = 1
    }
    await init()
}

const nextPage = async () => {
    page++
    await init()
}

const setPageSize = async (size) => {
    if (pageSize != size) {
        pageSize = size
        await init();
    }
}

const submitPageSize = async () => {
    const size = document.getElementById("pagesizeInput").value
    if (size) {
        await setPageSize(size)
    }
}