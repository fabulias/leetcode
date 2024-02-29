#include <iostream>
using namespace std;
template <typename T>
class Arbolito;

template <typename T>
class Nodo {
	friend class Arbolito<T>;
	protected:
		T valor_;
		Nodo<T> *left_;
		Nodo<T> *right_;
	public:
		Nodo() {
			left_ = NULL;
			right_= NULL;
		}
		Nodo(T valor) {
			left_ = NULL;
			right_= NULL;
			valor_ = valor;
		}

};

template <typename T>
class Arbolito {
	protected:
		Nodo<T> *raiz_;
	public:
		Arbolito() {
			raiz_ = NULL;
		}

		Nodo<T>* MinOfMax(Nodo<T> *raiz){
			if(raiz == NULL)
				return NULL;
			if(raiz->left_ == NULL)
				return raiz;
			else
				MinOfMax(raiz->left_);
		}
		void Insert(T valor, Nodo<T> *&n) {
			if (n == NULL) {
				Nodo<T> *nuevo = new Nodo<T>(valor);
				n = nuevo;
				return;
			}
			if (n->valor_ == valor) {
				return;
			}
			if (n->valor_ > valor) {
				Insert(valor,n->left_);
			} else {
				Insert(valor,n->right_);
			}
		}

		Nodo<T> *&GetRaiz() {
			return raiz_;
		}
		void Preorder(Nodo<T> *raiz) {
			if(raiz == NULL) {
				return;
			}
			cout << raiz->valor_ << endl;
			Inorder(raiz->left_);
			Inorder(raiz->right_);
		}		
		void Inorder(Nodo<T> *raiz) {
			if(raiz == NULL) {
				return;
			}
			Inorder(raiz->left_);
			cout << raiz->valor_ << endl;
			Inorder(raiz->right_);
		}
		void Postorder(Nodo<T> *raiz) {
			if(raiz == NULL) {
				return;
			}
			Inorder(raiz->left_);
			Inorder(raiz->right_);
			cout << raiz->valor_ << endl;
		}

		int Altura(Nodo<T> *ptr){
        if(ptr==NULL)
            return -1;
        else
            return (1+ max(Altura(ptr->left_),Altura(ptr->right_)));
    	}

    	int Size(Nodo<T> *ptr){
        if(ptr==NULL)
            return 0;
        else
            return(1+ Size(ptr->left_) + Size(ptr->right_));
    	}

    	int Search(Nodo<T> *ptr, T n){
        if(ptr->valor_ == n) return 1;
        else{
            if(ptr->getId() < n)
                Search(ptr->left_, n);
            else if(ptr->getId() >= n)
                Search(ptr->right_, n);
            return 0;
        }
    }

		void Borrar(Nodo<T> *&raiz, T valor){
			if(raiz==NULL)
				return;
			else if(valor < raiz->valor_){
				Borrar(raiz->left_,valor);
			}
			else if(valor > raiz->valor_)
				Borrar(raiz->right_,valor);
			else{
				if(raiz->left_ == NULL)
				{
					Nodo<T> *tmp = raiz;
					raiz = raiz->right_;
					delete tmp;
					return;
				}
				if(raiz->right_ == NULL){
					Nodo<T> *tmp = raiz;
					raiz = raiz->left_;
					delete tmp;
					return;
				}
				else{ 
					Nodo<T> *aux;
					aux = MinOfMax(raiz->right_);		
					raiz->valor_ = aux->valor_;
					Borrar(raiz->right_,raiz->valor_);
					aux = NULL;
					delete aux;
				}
			}
		}

};

int main() {
	Arbolito<int> *arbol = new Arbolito<int>();
	Nodo<int> *raiz = arbol->GetRaiz();
	arbol->Insert(10,raiz);
	arbol->Insert(5,raiz);
	arbol->Insert(20,raiz);
	arbol->Insert(15,raiz);
	arbol->Insert(25,raiz);
	arbol->Insert(0,raiz);
	arbol->Insert(30,raiz);
	arbol->Insert(8,raiz);
	arbol->Preorder(raiz);
	arbol->Borrar(raiz,25);
	cout<<endl;
	arbol->Inorder(raiz);
	arbol->Borrar(raiz,15);
	cout<<endl;
	arbol->Postorder(raiz);
	arbol->Borrar(raiz,1);
	cout<<endl;
	cout<<"Altura : "<<arbol->Altura(raiz)<<endl;
	cout<<"Cantidad nodos : "<<arbol->Size(raiz)<<endl;

}
