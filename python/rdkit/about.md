# rdkit in python

## [安装](http://www.rdkit.org/docs/Install.html)

    anaconda
        wget https://repo.continuum.io/archive/Anaconda3-5.1.0-Linux-x86_64.sh
        bash Anaconda3-5.1.0-Linux-x86_64.sh -u
        conda create -c rdkit -n my-rdkit-env rdkit
        source activate my-rdkit-env
    ubuntu
        sudo apt-get install python-rdkit librdkit1 rdkit-data

## 使用

    python3
        from rdkit import CHem
            出现错误 // ImportError: libXrender.so.1: cannot open shared object file: No such file or directory
            解决 sudo apt-get install libxrender1
        m = Chem.MolFromSmiles('Cc1ccccc1')
        m.GetNumAtoms()
        Chem.MolToSmiles(m)
        Chem.MolToMolBlock(m)
        for atom in m.GetAtoms():
            print(atom.GetAtomicNum())
        m.GetAtomWithIdx(0).GetSymbol()
        m.GetAtomWithIdx(0).IsInRing()
        m2 = Chem.MolFromSmiles('CCO')
        m2 = Chem.AddHs(m)
        Chem.MolToMolBlock(m2)
        m3 = Chem.RemoveHs(m2)
        Chem.MolToMolBlock(m3)
        patt = Chem.MolFromSmarts('ccO')
        m.HasSubstructMatch(patt)
        m.GetSubstructMatch(patt)
