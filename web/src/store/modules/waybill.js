export const SET_DOCUMENT_ID = 'SET_DOCUMENT_ID';

export const state = {
    documentReceiptId: null
}

export const getters = {
    documentReceiptId: state => state.documentReceiptId,
}

export const mutations = {
    SET_DOCUMENT_ID: (state, payload) => {
        state.documentReceiptId = payload;
    },
}

export const actions = {
    setDocumentId({commit}, payload) {
        commit(SET_DOCUMENT_ID, payload);
    },
}